package routers

import (
	h "github.com/atrariksa/golang-service-framework/src/handlers"
	s "github.com/atrariksa/golang-service-framework/src/server"
	services "github.com/atrariksa/golang-service-framework/src/services"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Server s.Server
	Logger *logrus.Logger
}

func SetUpRouter(server s.Server, logger *logrus.Logger) Router {
	router := Router{
		Server: server,
		Logger: logger,
	}
	return router
}

func (r Router) Route() {
	// set up simple api
	simpleApiService := services.SetUpSimpleApiService(r.Logger)
	simpleApiHandler := h.SetUpSimpleApiHandler(simpleApiService, r.Logger)
	r.Server.AddRoute("/simple", simpleApiHandler.HandleSimpleApi)

	// set up simple temporary value holder api
	valueHolderApiService := services.SetUpValueHolderApiService(r.Logger)
	valueHolderApiHandler := h.SetUpValueHolderApiHandler(valueHolderApiService, r.Logger)
	r.Server.AddRoute("/value-holder", valueHolderApiHandler.HandleValueHolder)
}
