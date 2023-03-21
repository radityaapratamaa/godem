package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"godem/domain/models"
	usermodel "godem/domain/models/user"
	"godem/infrastructure/common"
)

func TestMaster_GetList(t *testing.T) {
	initTest()
	mockGetListResponse := []*usermodel.Users{{ID: 1}}
	mockRequestData := &usermodel.UsersRequest{}
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := &models.SelectResponse{
			RequestParam: mockRequestData,
			Data:         mockGetListResponse,
		}
		masterRepo.On("GetList", mock.Anything, mock.Anything).
			Return(mockGetListResponse, nil).Once()
		actualResult, err := masterUc.GetList(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	t.Run("Case 2 - Failed", func(t *testing.T) {
		masterRepo.On("GetList", mock.Anything, mock.Anything).
			Return(nil, common.ErrPatch).Once()
		actualResult, err := masterUc.GetList(context.Background(), mockRequestData)
		assert.Equal(t, nil, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_GetDetailByID(t *testing.T) {
	initTest()
	mockRequestData := int64(1)
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := &usermodel.Users{ID: 1}
		masterRepo.On("GetDetailByID", mock.Anything, mock.Anything).
			Return(expectedResult, nil).Once()
		actualResult, err := masterUc.GetDetailByID(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	t.Run("Case 2 - Failed", func(t *testing.T) {
		masterRepo.On("GetDetailByID", mock.Anything, mock.Anything).
			Return(nil, common.ErrPatch).Once()
		actualResult, err := masterUc.GetDetailByID(context.Background(), mockRequestData)
		assert.Equal(t, nil, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_CreateNew(t *testing.T) {
	initTest()
	mockRequestData := &usermodel.Users{}
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := &models.CUDResponse{}
		masterRepo.On("CreateNew", mock.Anything, mock.Anything).
			Return(expectedResult, nil).Once()
		actualResult, err := masterUc.CreateNew(context.Background(), mockRequestData)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
}

func TestMaster_UpdateData(t *testing.T) {
	initTest()
	mockRequestData := &usermodel.Users{}
	mockID := int64(1)
	expectedResult := &models.CUDResponse{}
	t.Run("Case 1 - Success", func(t *testing.T) {
		masterRepo.On("UpdateData", mock.Anything, mock.Anything, mock.Anything).
			Return(expectedResult, nil).Once()
		actualResult, err := masterUc.UpdateData(context.Background(), mockRequestData, mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
	expectedResult = nil
	t.Run("Case 2 - Failed Parse Interface", func(t *testing.T) {
		newMockRequestData := &usermodel.LoginResponse{}
		actualResult, err := masterUc.UpdateData(context.Background(), newMockRequestData, mockID)
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, true, err != nil)
	})
}

func TestMaster_DeleteData(t *testing.T) {
	initTest()
	t.Run("Case 1 - Success", func(t *testing.T) {
		expectedResult := &models.CUDResponse{}
		masterRepo.On("DeleteData", mock.Anything, mock.Anything).
			Return(expectedResult, nil).Once()
		actualResult, err := masterUc.DeleteData(context.Background(), int64(1))
		assert.Equal(t, expectedResult, actualResult)
		assert.Equal(t, false, err != nil)
	})
}
