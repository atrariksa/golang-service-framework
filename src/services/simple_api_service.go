package services

import (
	c "github.com/atrariksa/golang-service-framework/src/constants"
	req "github.com/atrariksa/golang-service-framework/src/models/requests"
	res "github.com/atrariksa/golang-service-framework/src/models/responses"
	"github.com/sirupsen/logrus"
)

type SimpleApiService struct {
	Logger *logrus.Logger
}

func SetUpSimpleApiService(logger *logrus.Logger) SimpleApiService {
	service := SimpleApiService{
		Logger: logger,
	}
	return service
}

type ISimpleApiService interface {
	Respond(request req.SimpleApiRequest) res.SimpleApiResponse
}

func (s SimpleApiService) Respond(request req.SimpleApiRequest) res.SimpleApiResponse {
	var message string
	if request.Message != "" {
		message = c.RECEIVED_MESSAGE + request.Message
	}
	response := res.SimpleApiResponse{
		Status:  c.STATUS_OK,
		Message: message,
	}
	return response
}
