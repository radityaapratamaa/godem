package user

import (
	"encoding/json"
	"godem/lib/util/response"
	"net/http"
	"strconv"

	"github.com/gorilla/schema"

	"github.com/go-chi/chi/v5"

	usermodel "godem/domain/models/user"
	"godem/infrastructure/common"
	"godem/usecase/user"
)

type Masters interface {
	common.CRUDHandler
}

type Master struct {
	uc user.Masters
}

func NewMaster(ucMaster user.Masters) *Master {
	return &Master{uc: ucMaster}
}

func (m *Master) GetList(w http.ResponseWriter, r *http.Request) {
	var requestData usermodel.UsersRequest
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	r.ParseForm()
	if err := schema.NewDecoder().Decode(&requestData, r.URL.Query()); err != nil {
		resp.BadRequest(err)
		return
	}

	data, err := m.uc.GetList(r.Context(), &requestData)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(data)
	return
}

func (m *Master) GetDetailByID(w http.ResponseWriter, r *http.Request) {
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		resp.BadRequest(err)
		return
	}

	data, err := m.uc.GetDetailByID(r.Context(), id)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(data)
	return
}

func (m *Master) CreateNew(w http.ResponseWriter, r *http.Request) {
	var requestData usermodel.Users
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		resp.BadRequest(err)
		return
	}

	data, err := m.uc.CreateNew(r.Context(), &requestData)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(data)
	return
}

func (m *Master) UpdateData(w http.ResponseWriter, r *http.Request) {
	var requestData usermodel.Users
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		resp.BadRequest(err)
		return
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		resp.BadRequest(err)
		return
	}

	data, err := m.uc.UpdateData(r.Context(), &requestData, id)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(data)
	return
}

func (m *Master) DeleteData(w http.ResponseWriter, r *http.Request) {
	resp := response.NewJSON()
	defer func() {
		resp.Send(w)
	}()

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		resp.BadRequest(err)
		return
	}

	data, err := m.uc.DeleteData(r.Context(), id)
	if err != nil {
		resp.InternalServerError(err)
		return
	}

	resp.Success(data)
	return
}
