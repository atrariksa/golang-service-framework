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

var vHolderHeaders = map[string]string{c.CONTENT_TYPE: c.JSON_CONTENT_TYPE}

type ValueHolderApiHandler struct {
	ValueHolderApiService s.ValueHolderApiService
	Logger                *logrus.Logger
}

func SetUpValueHolderApiHandler(ValueHolderApiService s.ValueHolderApiService, logger *logrus.Logger) ValueHolderApiHandler {
	handler := ValueHolderApiHandler{
		ValueHolderApiService: ValueHolderApiService,
		Logger:                logger,
	}
	return handler
}

type IValueHolderApiHandler interface {
	HandleValueHolder(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

/**
This method handle incoming request to Value Holder Api.
Content-type for both request and response are application json.
*/
func (s ValueHolderApiHandler) HandleValueHolder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.Get(w, r)
		return
	} else if r.Method == http.MethodPost {
		s.Post(w, r)
		return
	} else {
		err := e.New(http.StatusText(http.StatusNotFound))
		s.Logger.Error(err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
}

func (s ValueHolderApiHandler) Get(w http.ResponseWriter, r *http.Request) {
	fields := make(map[string]interface{})
	var err e.Err
	id := r.URL.Query()[c.ID]
	if len(id) < 1 {
		err = e.New(http.StatusText(http.StatusBadRequest))
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	resp, err := s.ValueHolderApiService.GetObject(id[0])
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

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
}
func (s ValueHolderApiHandler) Post(w http.ResponseWriter, r *http.Request) {
	fields := make(map[string]interface{})

	reqStruct := req.ValueHolderApiRequest{}
	err := val.ValidateHeader(r, vHolderHeaders)
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

	id := r.URL.Query()[c.ID]
	err = s.ValueHolderApiService.CreateObject(id[0], reqStruct)
	if err != nil {
		s.Logger.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
