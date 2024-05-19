package user

import (
	"context"
	"embed"
	"fmt"
	common2 "godem/src/common"
	user2 "godem/src/domain/models/user"
	"godem/src/lib/helper"
	userdb "godem/src/repository/user"
	"os"
	"time"

	"github.com/kodekoding/phastos/v2/go/api"
	"github.com/kodekoding/phastos/v2/go/cache"
	"github.com/kodekoding/phastos/v2/go/database"
	helperphastos "github.com/kodekoding/phastos/v2/go/helper"
	"github.com/kodekoding/phastos/v2/go/mail"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/null"
)

type Auths interface {
	Authenticate(ctx context.Context, requestData *user2.LoginRequest) (*user2.LoginResponse, error)
	ForgotPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, requestData *user2.ResetPasswordRequest) error
	UpdatePassword(ctx context.Context, requestData *user2.ChangePasswordRequest) error
	ResendToken(ctx context.Context, email string) error
	VerifyToken(ctx context.Context, email, token string) error
	Logout(ctx context.Context) (*database.CUDResponse, error)
}

type auth struct {
	smtp           mail.SMTPs
	repo           userdb.Auths
	cache          cache.Caches
	templateFolder embed.FS
}

var (
	generateJWTToken = helperphastos.GenerateJWTToken
)

func newAuth(
	repo userdb.Auths,
	smtp mail.SMTPs,
	template embed.FS,
	cache cache.Caches,
) *auth {
	return &auth{repo: repo, smtp: smtp, templateFolder: template, cache: cache}
}

func (l *auth) Authenticate(ctx context.Context, requestData *user2.LoginRequest) (*user2.LoginResponse, error) {
	data, err := l.repo.Authenticate(ctx, requestData)
	if err != nil {
		return nil, api.NotFound("user not found", "USER_NOT_FOUND")
	}

	savedDeviceId := data.DeviceId.String
	requestedDeviceId := requestData.DeviceId.String

	if !data.DeviceId.Valid {
		// async process to update the device id
		go func() {
			updatingData := new(user2.Data)
			updatingData.DeviceId = requestData.DeviceId
			if _, err = l.repo.UpdateById(context.Background(), updatingData, data.Id); err != nil {
				log.Error().Msgf("Failed to update device id data: %s", err.Error())

				_ = helperphastos.SendSlackNotification(context.Background(),
					helperphastos.NotifMsgType(helperphastos.NotifErrorType),
					helperphastos.NotifChannel(os.Getenv("NOTIFICATIONS_SLACK_WEBHOOK_URL")),
					helperphastos.NotifTitle("Failed to Update Device ID"),
					helperphastos.NotifData(map[string]string{
						"description": err.Error(),
						"email":       requestData.Email,
						"device_id":   requestData.DeviceId.String,
					}),
				)
			}
		}()
	} else if savedDeviceId != requestedDeviceId && requestedDeviceId != "" {
		return nil, api.Unauthorized("unauthorized device", "UNAUTHORIZED_DEVICE")
	}

	if data.FirebaseId.String != requestData.FirebaseId.String && requestData.FirebaseId.String != "" {
		// async process to update the device id
		go func() {
			updatingData := new(user2.Data)
			updatingData.FirebaseId = requestData.FirebaseId
			if _, err = l.repo.UpdateById(context.Background(), updatingData, data.Id); err != nil {
				log.Error().Msgf("Failed to update firebase id data: %s", err.Error())
				_ = helperphastos.SendSlackNotification(context.Background(),
					helperphastos.NotifMsgType(helperphastos.NotifErrorType),
					helperphastos.NotifChannel(os.Getenv("NOTIFICATIONS_SLACK_WEBHOOK_URL")),
					helperphastos.NotifTitle("Failed to Update Firebase ID"),
					helperphastos.NotifData(map[string]string{
						"description": err.Error(),
						"email":       requestData.Email,
						"firebase_id": requestData.FirebaseId.String,
					}),
				)
			}
		}()
	}

	var jwtData user2.JWTData
	tableReq := new(database.TableRequest)
	tableReq.SetWhereCondition("user_id = ?", data.Id)
	if err = l.repo.GetDetail(ctx, &database.QueryOpts{
		SelectRequest:     tableReq,
		OptionalTableName: "vw_employee_profile",
		Columns:           "id, user_id, email, account_id, manager_id, employee_id",
		Result:            &jwtData,
	}); err != nil {
		return nil, errors.Wrap(err, "usecase.user.auth.Authenticate.GetDetailEmployeeProfile")
	}

	if !jwtData.EmployeeId.Valid {
		return nil, api.Forbidden("user is exists, but not registered as employee", "NON_EMPLOYEE_USER")
	}

	jwtToken, err := generateJWTToken(jwtData)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.auth.Authenticate.GenerateJWTToken")
	}

	response := new(user2.LoginResponse)
	response.Token = jwtToken
	return response, nil
}

func (l *auth) ForgotPassword(ctx context.Context, email string) error {
	tableReq := new(database.TableRequest)
	tableReq.SetWhereCondition("email = ?", email)
	var result user2.Data
	if err := l.repo.GetDetail(ctx, &database.QueryOpts{
		SelectRequest: tableReq,
		Columns:       "email, name, activation_code, activation_expired_at",
		Result:        &result,
	}); err != nil {
		return errors.Wrap(err, "usecase.user.auth.ForgotPassword.CheckEmail")
	}

	if result.Email == "" {
		return api.NotFound("email not found", "EMAIL_NOT_FOUND")
	}

	timeNow := time.Now()
	if result.ActivationCode.Valid {
		getExpiredAt, _ := time.Parse(time.RFC3339, result.ActivationExpiredAt.String)
		diff := getExpiredAt.Sub(timeNow)
		if diff.Seconds() >= 0 {
			return api.BadRequest("user has already requested the password change, please check your email or wait for 24 hours to re-request", "ALREADY_REQUESTED")
		}
	}

	activationToken := helperphastos.GenerateRandomStringWithCharset(6, "0123456789")
	expiredAt := timeNow.Add(24 * time.Hour)

	users := &user2.Data{
		ActivationCode:      null.StringFrom(activationToken),
		ActivationExpiredAt: null.StringFrom(expiredAt.Format("2006-01-02 15:04:05")),
		ActiveAt:            null.StringFrom("null"),
	}
	if _, err := l.repo.Update(ctx, users, map[string]interface{}{
		"email = ?": email,
	}); err != nil {
		return errors.Wrap(err, "usecase.user.auth.ForgotPassword.UpdateActivationCode")
	}

	go func() {
		type forgotPassArgs struct {
			FullName string
			OTPCode  string
		}

		args := &forgotPassArgs{
			FullName: result.Name,
			OTPCode:  activationToken,
		}
		if err := l.smtp.
			AddRecipient(result.Email).
			SetHTMLTemplate(l.templateFolder, common2.EmailTemplate[common2.ForgotPassword], "Request To Change Password", args).
			Send(); err != nil {
			log.Error().Msg("Error When Sent Email to " + result.Email + ": " + err.Error())
		}
	}()

	return nil
}

func (l *auth) ResendToken(ctx context.Context, email string) error {
	cacheKey := fmt.Sprintf("%s:otp-request:%s", common2.PrefixCacheKey, email)
	value, err := l.cache.Get(ctx, cacheKey)
	if err == nil && value == "ok" {
		return api.BadRequest("user already request the otp, please wait 1 minute to re-request the otp", "USER_ALREADY_REQUEST_OTP")
	}
	if _, err = l.cache.Set(ctx, cacheKey, "ok", 60); err != nil {
		return errors.Wrap(err, "usecase.user.auth.ResendToken.SetCache")
	}

	return l.ForgotPassword(ctx, email)
}

func (l *auth) VerifyToken(ctx context.Context, email, token string) error {
	tableReq := new(database.TableRequest)
	tableReq.SetWhereCondition("email = ?", email)
	tableReq.SetWhereCondition("activation_code = ?", token)
	var result user2.Data
	if err := l.repo.GetDetail(ctx, &database.QueryOpts{
		SelectRequest: tableReq,
		Columns:       "activation_expired_at",
		Result:        &result,
	}); err != nil {
		return errors.Wrap(err, "usecase.user.auth.ResetPassword.CheckEmailxOTP")
	}

	if !result.ActivationExpiredAt.Valid {
		return api.Forbidden("not authorized user", "FORBIDDEN_ACCESS")
	}
	timeNow := time.Now()
	if result.ActivationExpiredAt.Valid {
		getExpiredAt, _ := time.Parse(time.RFC3339, result.ActivationExpiredAt.String)
		diff := getExpiredAt.Sub(timeNow)
		if diff.Seconds() < 0 {
			return api.BadRequest("activation code was expired", "EXPIRED_ACTIVATION_CODE")
		}
	}
	return nil
}

func (l *auth) ResetPassword(ctx context.Context, requestData *user2.ResetPasswordRequest) error {
	if err := l.VerifyToken(ctx, requestData.Email, requestData.OTP); err != nil {
		return err
	}

	if requestData.NewPassword != requestData.ConfirmNewPassword {
		return api.BadRequest("new password doesn't match", "CONFIRM_PASSWORD_NOT_MATCH")
	}

	if err := l.repo.ResetPassword(ctx, requestData); err != nil {
		return errors.Wrap(err, "usecase.user.auth.ResetPassword.ExecDB")
	}

	return nil
}

func (l *auth) UpdatePassword(ctx context.Context, requestData *user2.ChangePasswordRequest) error {
	if requestData.NewPassword != requestData.ConfirmNewPassword {
		return api.BadRequest("new password doesn't match", "CONFIRM_PASSWORD_NOT_MATCH")
	}

	jwtData := helper.GetJWTData(ctx)
	requestData.Email = jwtData.Email
	if _, err := l.repo.Authenticate(ctx, &user2.LoginRequest{
		Email:    jwtData.Email,
		Password: null.StringFrom(requestData.OldPassword),
	}); err != nil {
		return api.BadRequest("not valid password", "NOT_VALID_PASSWORD")
	}

	if err := l.repo.UpdatePassword(ctx, requestData); err != nil {
		return errors.Wrap(err, "usecase.user.auth.UpdatePassword.ExecDB")
	}

	return nil
}

func (l *auth) Logout(ctx context.Context) (*database.CUDResponse, error) {
	jwtData := helper.GetJWTData(ctx)
	willBeUpdated := new(user2.Data)
	willBeUpdated.DeviceId = null.StringFrom("null")
	exec, err := l.repo.UpdateById(ctx, willBeUpdated, jwtData.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.ResetDeviceId.UpdateById")
	}
	return exec, nil
}
