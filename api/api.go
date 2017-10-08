package api

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"

	"git.icysoft.fr/cedric/industry-go-server/service"
	"github.com/xeipuuv/gojsonschema"
)

var playerService *service.PlayerService

func init() {
	playerService = &service.PlayerService{}
}

func validateData(schema *gojsonschema.Schema, body io.ReadCloser) (bool, io.ReadCloser, error) {
	var bodyBytes []byte
	if body != nil {
		bodyBytes, _ = ioutil.ReadAll(body)
	}

	// Restore the io.ReadCloser to its original state
	body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// Use the content
	s := string(bodyBytes)

	validation, err := newPlayerSchema.Validate(gojsonschema.NewStringLoader(s))
	if err != nil {
		return false, body, err
	}
	if !validation.Valid() {
		validationErrorStr := "Data validation error : "
		for _, validationErr := range validation.Errors() {
			validationErrorStr += validationErr.String()
		}
		return false, body, errors.New(validationErrorStr)
	}
	return true, body, nil
}
