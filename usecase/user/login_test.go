package user

import (
	"context"
	usermodel "godem/domain/models/user"
	"godem/infrastructure/common"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestLogin_Authenticate(t *testing.T) {
	initTest()
	expectedResult := &usermodel.LoginResponse{
		Token:   "jwtToken",
		Expired: 1,
	}
	mockAuthenticateResult := &usermodel.Users{}
	mockRequestData := &usermodel.LoginRequest{}
	t.Run("Case 1 - Success", func(t *testing.T) {
		login.On("Authenticate", mock.Anything, mock.Anything).
			Return(mockAuthenticateResult, nil).Once()
		generateJWTToken = func(signingKey string, data interface{}) (string, int64, error) {
			return expectedResult.Token, expectedResult.Expired, nil
		}

		actualResult, err := uc.login.Authenticate(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})

	expectedResult = nil
	t.Run("Case 2 - Failed GenerateJWTToken", func(t *testing.T) {
		login.On("Authenticate", mock.Anything, mock.Anything).
			Return(mockAuthenticateResult, nil).Once()
		generateJWTToken = func(signingKey string, data interface{}) (string, int64, error) {
			return "", 0, common.ErrPatch
		}

		actualResult, err := uc.login.Authenticate(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
	t.Run("Case 3 - Failed Authenticate to DB", func(t *testing.T) {
		login.On("Authenticate", mock.Anything, mock.Anything).
			Return(nil, common.ErrPatch).Once()

		actualResult, err := uc.login.Authenticate(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
	t.Run("Case 4 - Failed Decode base64 string", func(t *testing.T) {
		decodeBase64String = func(s string) ([]byte, error) {
			return nil, common.ErrPatch
		}
		actualResult, err := uc.login.Authenticate(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}
