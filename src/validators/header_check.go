package validators

import (
	"net/http"

	c "github.com/atrariksa/golang-service-framework/src/constants"
	e "github.com/atrariksa/golang-service-framework/src/errors"
)

type RequestHeaderValidator struct {
}

type IRequestHeaderValidator interface {
	ValidateHeader(r *http.Request, headers map[string]string) e.Err
}

func ValidateHeader(r *http.Request, headers map[string]string) e.Err {
	var valid bool
	for k, v := range headers {
		if r.Header.Get(k) == v {
			valid = true
		} else {
			valid = false
			break
		}
	}
	if valid != true {
		err := e.New(c.INVALID_HEADER)
		return err
	} else {
		return nil
	}
}
