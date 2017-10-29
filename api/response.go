package api

import (
	"encoding/json"

	"github.com/cmillauriaux/industry-go-server/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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

func convertStructToAnother(source interface{}, destination interface{}) interface{} {
	copier.Copy(destination, source)
	return destination
}

func makeResponse(c *gin.Context, source interface{}, destination interface{}) {
	data := convertStructToAnother(source, destination)
	response := structToJSON(Response{Data: data})
	c.JSON(200, response)
}

func makeError(err error) *Response {
	businessErr, ok := err.(*model.BusinessError)
	if ok {
		return &Response{Code: businessErr.Code, Err: businessErr.Error()}
	}
	return &Response{Code: -1, Err: err.Error()}
}
