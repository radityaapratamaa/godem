package user

import (
	"context"
	"godem/domain/models"
	"godem/domain/models/user"
	"godem/infrastructure/common"
	"godem/infrastructure/database/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMaster_GetList(t *testing.T) {
	initTest()

	mockRequest := &user.UsersRequest{Query: StrGetContext}
	var mockParams []interface{}
	for i := 0; i < 6; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := []*user.Users{
			{ID: 1},
		}
		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrSelectContext, mockParams...).
			Return(func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
				data := dest.(*[]*user.Users)
				*data = expectedResult
				return nil
			}).Once()

		actualResult, err := masterUser.GetList(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrSelectContext, mockParams...).
			Return(common.ErrPatch).Once()

		actualResult, err := masterUser.GetList(context.Background(), mockRequest)
		assert.Equal(t, nil, actualResult)
		assert.Equal(t, true, err != nil)
	})

	t.Run("Case 3 - Failed Parse Interface", func(t *testing.T) {
		newMockReq := &user.LoginRequest{}
		actualResult, err := masterUser.GetList(context.Background(), newMockReq)
		assert.Equal(t, nil, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_GetDetailByID(t *testing.T) {
	initTest()
	mockRequest := int64(1)
	var mockParams []interface{}
	for i := 0; i < 4; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := &user.Users{
			ID: 1,
		}

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrGetContext, mockParams...).
			Return(func(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
				data := dest.(*user.Users)
				*data = *expectedResult
				return nil
			}).Once()

		actualResult, err := masterUser.GetDetailByID(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		followerRepo.On(StrGetContext, mockParams...).
			Return(common.ErrPatch).Once()

		actualResult, err := masterUser.GetDetailByID(context.Background(), mockRequest)
		assert.Equal(t, nil, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_CreateNew(t *testing.T) {
	initTest()
	mockRequest := &user.Users{}
	var mockParams []interface{}
	for i := 0; i < 9; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	mockSQLResult := mocks.SQLResult{}
	expectedResult := &models.CUDResponse{
		Status:       true,
		RowsAffected: 1,
		LastInsertID: 1,
	}
	t.Run("Case 1 - Success", func(t *testing.T) {
		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(mockSQLResult, nil).Once()

		actualResult, err := masterUser.CreateNew(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	expectedResult = nil
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(nil, common.ErrPatch).Once()

		actualResult, err := masterUser.CreateNew(context.Background(), mockRequest)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})

	t.Run("Case 3 - Failed Parse Interface", func(t *testing.T) {
		newMockReq := &user.LoginRequest{}
		actualResult, err := masterUser.CreateNew(context.Background(), newMockReq)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_UpdateData(t *testing.T) {
	initTest()
	mockRequest := &user.Users{}
	mockID := int64(1)
	var mockParams []interface{}
	for i := 0; i < 9; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	mockSQLResult := mocks.SQLResult{}
	expectedResult := &models.CUDResponse{
		Status:       true,
		RowsAffected: 1,
		LastInsertID: 1,
	}
	t.Run("Case 1 - Success", func(t *testing.T) {
		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(mockSQLResult, nil).Once()

		actualResult, err := masterUser.UpdateData(context.Background(), mockRequest, mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	expectedResult = nil
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(nil, common.ErrPatch).Once()

		actualResult, err := masterUser.UpdateData(context.Background(), mockRequest, mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
	t.Run("Case 3 - Failed Parse Interface", func(t *testing.T) {
		newMockReq := &user.LoginRequest{}
		actualResult, err := masterUser.UpdateData(context.Background(), newMockReq, mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_DeleteData(t *testing.T) {
	initTest()
	mockID := int64(1)
	var mockParams []interface{}
	for i := 0; i < 3; i++ {
		mockParams = append(mockParams, mock.Anything)
	}
	mockSQLResult := mocks.SQLResult{}
	expectedResult := &models.CUDResponse{
		Status:       true,
		RowsAffected: 1,
		LastInsertID: 1,
	}
	t.Run("Case 1 - Success", func(t *testing.T) {
		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(mockSQLResult, nil).Once()

		actualResult, err := masterUser.DeleteData(context.Background(), mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	expectedResult = nil
	t.Run("Case 2 - Failed", func(t *testing.T) {

		masterRepo.On("Rebind", mock.Anything).Return("").Once()
		masterRepo.On(StrExecContext, mockParams...).
			Return(nil, common.ErrPatch).Once()

		actualResult, err := masterUser.DeleteData(context.Background(), mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}
