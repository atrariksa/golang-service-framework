package services

import (
	c "github.com/atrariksa/golang-service-framework/src/constants"
	e "github.com/atrariksa/golang-service-framework/src/errors"
	req "github.com/atrariksa/golang-service-framework/src/models/requests"
	res "github.com/atrariksa/golang-service-framework/src/models/responses"
	"github.com/sirupsen/logrus"
)

var maxCache = 3
var Cache = make(map[string]res.ValueHolderApiResponse, maxCache)

type ValueHolderApiService struct {
	Logger *logrus.Logger
}

func SetUpValueHolderApiService(logger *logrus.Logger) ValueHolderApiService {
	service := ValueHolderApiService{
		Logger: logger,
	}
	return service
}

type IValueHolderApiService interface {
	CreateObject(id string, request req.ValueHolderApiRequest) e.Err
	GetObject(id string) (res.ValueHolderApiResponse, e.Err)
}

func (s ValueHolderApiService) CreateObject(id string, request req.ValueHolderApiRequest) e.Err {
	if _, ok := Cache[id]; ok {
		err := e.New(c.DATA_ALREADY_ADDED)
		return err
	}
	if len(Cache) < maxCache {
		Cache[id] = res.ValueHolderApiResponse{
			ID:   id,
			Name: request.Name,
			Role: request.Role,
		}
		return nil
	} else {
		err := e.New(c.CACHE_IS_FULL)
		return err
	}
}

func (s ValueHolderApiService) GetObject(id string) (res.ValueHolderApiResponse, e.Err) {
	if response, ok := Cache[id]; ok {
		return response, nil
	} else {
		err := e.New(c.VALUE_NOT_FOUND)
		return response, err
	}
}
