package utils

import (
	"encoding/json"
	"net/http"

	"github.com/alka/supermartchannel/api"
)

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	if response != nil {
		json.NewEncoder(rw).Encode(response)
	}
}

func WriteErrorResponse(status int, err error, rw http.ResponseWriter) {
	martErr, ok := err.(*api.MartError)
	if !ok {
		martErr = &api.MartError{
			Code:        0,
			Message:     "failed in serving request",
			Description: martErr.Error(),
		}
	} else {
		status = martErr.Code.HTTPStatus()
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(martErr)
}
