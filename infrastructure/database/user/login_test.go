package user

import (
	"context"
	"godem/domain/models/user"
	"godem/infrastructure/common"
	"godem/infrastructure/database/mocks"
	"godem/lib/util/database"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

var (
	masterRepo   *mocks.Master
	followerRepo *mocks.Follower
	db           *database.DB
	login        *Login
	master       *Master
)

const (
	StrGetContext    = "GetContext"
	StrSelectContext = "SelectContext"
	StrExecContext   = "ExecContext"
)

func initTest() {
	masterRepo = new(mocks.Master)
	followerRepo = new(mocks.Follower)
	db = &database.DB{
		Master:   masterRepo,
		Follower: followerRepo,
	}
	login = NewLogin(db)
	master = NewMaster(db)
}

func TestLogin_Authenticate(t *testing.T) {
	initTest()
	expectedResult := &user.Users{ID: 1}
	mockRequest := &user.LoginRequest{}
	var mockParams []interface{}
	for i := 0; i < 5; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	t.Run("Case 1 - Success", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrGetContext, mockParams...).
			Return(func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
				data := dest.(*user.Users)
				*data = *expectedResult
				return nil
			}).Once()

		actualResult, err := login.Authenticate(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	expectedResult = nil
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrGetContext, mockParams...).
			Return(common.ErrPatch).Once()

		actualResult, err := login.Authenticate(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}
