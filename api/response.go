package api

import (
	"encoding/json"

	"git.icysoft.fr/cedric/industry-go-server/model"
)

type Response struct {
	Code int
	Err  string
	Data interface{}
}

func structToJSON(structure interface{}) string {
	b, err := json.Marshal(structure)
	if err != nil {
		return ""
	}
	return string(b)
}

func makeError(err error) *Response {
	businessErr, ok := err.(*model.BusinessError)
	if ok {
		return &Response{Code: businessErr.Code, Err: businessErr.Error()}
	}
	return &Response{Code: -1, Err: err.Error()}
}
