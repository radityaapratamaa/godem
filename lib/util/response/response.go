package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

type JSON struct {
	Code         int         `json:"api_code"`
	Message      string      `json:"api_display_message"`
	Error        error       `json:"-"`
	RawMessage   string      `json:"api_raw_message,omitempty"`
	Data         interface{} `json:"api_result_data"`
	Latency      float64     `json:"latency"`
	SlackTraceID string      `json:"slack_trace_id,omitempty"`
}

const (
	errMessageTemplate  = "%s : %s"
	ForbiddenResource   = "you're not authorized to access this features"
	InternalServerError = "Internal Server Error"
)

func NewJSON() *JSON {
	return &JSON{}
}
func (jr *JSON) SetMessage(msg string) *JSON {
	jr.Message = msg
	return jr
}

func (jr *JSON) SetData(data interface{}) *JSON {
	jr.Data = data
	jr.Code = 200
	return jr
}

func (jr *JSON) SetError(errCode int, err error) *JSON {
	jr.Code = errCode
	jr.Error = err
	return jr
}

func (jr *JSON) SetRawMessage(msg string) *JSON {
	jr.RawMessage = msg
	return jr
}

func (jr *JSON) Success(data interface{}) *JSON {
	jr.Data = data
	jr.Message = "SUCCESS"
	jr.Code = 200
	return jr
}

func (jr *JSON) BadRequest(err error) *JSON {
	jr.Message = "Bad Request"
	jr.RawMessage = err.Error()
	jr.Code = 400
	return jr
}

func (jr *JSON) ForbiddenResource(err error) *JSON {
	jr.Message = "Forbidden Resource"
	jr.RawMessage = fmt.Sprintf(errMessageTemplate, ForbiddenResource, errors.Cause(err).Error())
	jr.Code = 403
	jr.Error = err
	return jr
}

func (jr *JSON) Unauthorized(err error) *JSON {
	jr.Message = "Unauthorized"
	jr.RawMessage = fmt.Sprintf(errMessageTemplate, ForbiddenResource, errors.Cause(err).Error())
	jr.Code = 401
	jr.Error = err
	return jr
}

func (jr *JSON) InternalServerError(err error) *JSON {
	jr.Message = "Internal Server Error"
	jr.RawMessage = fmt.Sprintf(errMessageTemplate, InternalServerError, errors.Cause(err).Error())
	jr.Code = 500
	jr.Error = err
	return jr
}

func (jr *JSON) Send(w http.ResponseWriter) {
	if jr.Error != nil {
		log.Println("got an error: ", jr.Error.Error())
	}
	/* #nosec */
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(jr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, writeErr := w.Write([]byte(`{"errors":["Internal Server Error"]}`))
		if writeErr != nil {
			log.Println("error when write: ", writeErr.Error())
		}
	}

	w.WriteHeader(jr.Code)
	w.Write(b)
}
