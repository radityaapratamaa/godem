package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/kodekoding/phastos/go/common"
	"github.com/kodekoding/phastos/go/response"

	usermodel "godem/domain/models/user"
	"godem/usecase/user"
)

type Masters interface {
	common.HandlerCRUD
}

type master struct {
	uc user.Masters
}

func newMaster(ucMaster user.Masters) *master {
	return &master{uc: ucMaster}
}

func (m *master) GetList(w http.ResponseWriter, r *http.Request) *response.JSON {
	var requestData usermodel.UsersRequest

	r.ParseForm()
	if err := schema.NewDecoder().Decode(&requestData, r.URL.Query()); err != nil {
		return response.NewJSON().BadRequest(err)
	}

	data, err := m.uc.GetList(r.Context(), &requestData)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(data)
}

func (m *master) GetDetailById(w http.ResponseWriter, r *http.Request) *response.JSON {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return response.NewJSON().BadRequest(err)
	}

	data, err := m.uc.GetDetailByID(r.Context(), id)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(data)
}

func (m *master) Insert(w http.ResponseWriter, r *http.Request) *response.JSON {
	var requestData usermodel.Users
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return response.NewJSON().BadRequest(err)
	}

	data, err := m.uc.CreateNew(r.Context(), &requestData)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(data)
}

func (m *master) Update(w http.ResponseWriter, r *http.Request) *response.JSON {
	var requestData usermodel.Users

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		return response.NewJSON().BadRequest(err)
	}

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return response.NewJSON().BadRequest(err)
	}

	data, err := m.uc.UpdateData(r.Context(), &requestData, id)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(data)
}

func (m *master) Delete(w http.ResponseWriter, r *http.Request) *response.JSON {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		return response.NewJSON().BadRequest(err)
	}

	data, err := m.uc.DeleteData(r.Context(), id)
	if err != nil {
		return response.NewJSON().InternalServerError(err)
	}

	return response.NewJSON().Success(data)
}
