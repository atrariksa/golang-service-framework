package validators

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	c "github.com/atrariksa/golang-service-framework/src/constants"
	e "github.com/atrariksa/golang-service-framework/src/errors"
)

type RequestBodyValidator struct {
}

type IRequestBodyValidator interface {
	ValidateBody(r *http.Request, inputStruct interface{}) (string, e.Err)
}

func ValidateBody(r *http.Request, inputStruct interface{}) (string, e.Err) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = e.New(c.INVALID_BODY)
		return "", err
	}
	if err = json.Unmarshal(b, &inputStruct); err != nil {
		err = e.New(c.INVALID_JSON)
		return "", err
	}
	return string(b), nil
}
