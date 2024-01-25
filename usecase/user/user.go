package user

import (
	"context"
	"github.com/kodekoding/phastos/v2/go/common"
	"github.com/kodekoding/phastos/v2/go/database"
	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	employee2 "godem/domain/models/employee"
	usermodel "godem/domain/models/user"
	"godem/infrastructure/database/user"
	"godem/lib/helper"
)

type Masters interface {
	common.UsecaseCRUD
	ResetDeviceId(ctx context.Context, userId int64) (*database.CUDResponse, error)
	Me(ctx context.Context) (*database.SelectResponse, error)
}

type master struct {
	repo user.Masters
	trx  database.Transactions
}

func (m *master) GetList(ctx context.Context, requestData interface{}) (*database.SelectResponse, error) {
	reqData, valid := requestData.(*usermodel.Request)
	if !valid {
		return nil, errors.Wrap(common.ErrStructNotCompatible, "usecase.user.master.GetList.CastInterfaceToStruct")
	}

	var dataList []*usermodel.Data
	if err := m.repo.GetList(ctx, &database.QueryOpts{
		SelectRequest:  &reqData.TableRequest,
		Result:         &dataList,
		ExcludeColumns: "password",
	}); err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.GetList.GetDataFromDB")
	}

	totalData, totalFiltered, err := m.repo.Count(ctx, &reqData.TableRequest)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.GetList.CountFromDB")
	}

	response := &database.SelectResponse{
		ResponseMetaData: &database.ResponseMetaData{
			RequestParam:  requestData,
			TotalData:     int64(totalData),
			TotalFiltered: int64(totalFiltered),
		},
		Data: dataList,
	}

	return response, nil
}

func (m *master) GetDetailById(ctx context.Context, id interface{}) (*database.SelectResponse, error) {
	idInt, valid := id.(int)
	if !valid {
		return nil, errors.Wrap(common.ErrStructNotCompatible, "usecase.user.master.GetDetailById.CastIdInterface")
	}
	var userDetail usermodel.Data
	tableReq := new(database.TableRequest)
	tableReq.SetWhereCondition("id = ?", idInt)
	if err := m.repo.GetDetail(ctx, &database.QueryOpts{
		SelectRequest:  tableReq,
		Result:         &userDetail,
		ExcludeColumns: "password",
	}); err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.GetDetailById.GetDetailFromDB")
	}

	result := new(database.SelectResponse)
	result.Data = userDetail
	return result, nil
}

func (m *master) Insert(ctx context.Context, data interface{}) (*database.CUDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *master) Update(ctx context.Context, data interface{}) (*database.CUDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *master) Delete(ctx context.Context, id interface{}) (*database.CUDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *master) Me(ctx context.Context) (*database.SelectResponse, error) {
	jwtData := helper.GetJWTData(ctx)
	tableReq := new(database.TableRequest)
	tableReq.SetWhereCondition("user_id = ?", jwtData.UserId)
	var profileData employee2.Profile
	if err := m.repo.GetDetail(ctx, &database.QueryOpts{
		SelectRequest:     tableReq,
		Result:            &profileData,
		OptionalTableName: "vw_employee_profile",
	}); err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.me.GetDetailData")
	}

	result := new(database.SelectResponse)
	result.Data = profileData
	return result, nil
}

func (m *master) ResetDeviceId(ctx context.Context, userId int64) (*database.CUDResponse, error) {
	willBeUpdated := new(usermodel.Data)
	willBeUpdated.DeviceId = null.StringFrom("null")
	exec, err := m.repo.UpdateById(ctx, willBeUpdated, userId)
	if err != nil {
		return nil, errors.Wrap(err, "usecase.user.master.ResetDeviceId.UpdateById")
	}
	return exec, nil
}

func newMaster(
	repo user.Masters,
	trx database.Transactions,
) *master {
	return &master{
		repo: repo,
		trx:  trx,
	}
}
