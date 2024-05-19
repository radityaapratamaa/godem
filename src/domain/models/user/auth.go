package user

import "github.com/volatiletech/null"

type LoginRequest struct {
	Email      string      `json:"email" db:"email" validate:"required"`
	Password   null.String `json:"password,omitempty" db:"password" validate:"required"`
	DeviceId   null.String `json:"device_id" db:"device_id"`
	FirebaseId null.String `json:"firebase_id" db:"firebase_id"`
}

type LoginResponse struct {
	Token   string `json:"token"`
	Expired int64  `json:"expired"`
}

type ChangePasswordRequest struct {
	Email              string `json:"email"`
	OldPassword        string `json:"old_password" validate:"required"`
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}

type ResetPasswordRequest struct {
	VerifyTokenRequest
	NewPassword        string `json:"new_password" validate:"required"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}

type VerifyTokenRequest struct {
	OTP   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type JWTData struct {
	UserId                 int64       `json:"user_id" db:"user_id"`
	ManagerId              null.String `json:"manager_id" db:"manager_id"`
	AccountId              int64       `json:"account_id" db:"account_id"`
	EmployeeId             null.Int64  `json:"employee_id" db:"id"`
	NIK                    string      `json:"nik" db:"employee_id"`
	Email                  string      `json:"email" db:"email"`
	PlacementId            int64       `json:"placement_id"`
	DefaultShiftLocationId int64       `json:"default_shift_location_id"`
}
