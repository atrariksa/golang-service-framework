package handler

import (
	"encoding/json"
	"net/http"

	c "github.com/atrariksa/golang-service-framework/src/constants"
	e "github.com/atrariksa/golang-service-framework/src/errors"
	req "github.com/atrariksa/golang-service-framework/src/models/requests"
	s "github.com/atrariksa/golang-service-framework/src/services"
	val "github.com/atrariksa/golang-service-framework/src/validators"
	"github.com/sirupsen/logrus"
)

var sApiHeaders = map[string]string{c.CONTENT_TYPE: c.JSON_CONTENT_TYPE}

type SimpleApiHandler struct {
	SimpleApiService s.SimpleApiService
	Logger           *logrus.Logger
}

func SetUpSimpleApiHandler(simpleApiService s.SimpleApiService, logger *logrus.Logger) SimpleApiHandler {
	handler := SimpleApiHandler{
		SimpleApiService: simpleApiService,
		Logger:           logger,
	}
	return handler
}

type ISimpleApiHandler interface {
	HandleSimpleApi(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

/**
This method handle incoming request to Simple Api.
Content-type for both request and response are application json.
*/
func (s SimpleApiHandler) HandleSimpleApi(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		s.Post(w, r)
		return
	}
	err := e.New(http.StatusText(http.StatusNotFound))
	s.Logger.Error(err)
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(err.Error()))
	return
}

func (s SimpleApiHandler) Post(w http.ResponseWriter, r *http.Request) {
	fields := make(map[string]interface{})
	reqStruct := req.SimpleApiRequest{}

	err := val.ValidateHeader(r, sApiHeaders)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := val.ValidateBody(r, &reqStruct)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fields[c.REQ_BODY] = body

	resp := s.SimpleApiService.Respond(reqStruct)
	respBody, err := json.Marshal(resp)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fields[c.RES_BODY] = string(respBody)

	w.Header().Add(c.CONTENT_TYPE, c.JSON_CONTENT_TYPE)
	w.Write(respBody)
	return
}
