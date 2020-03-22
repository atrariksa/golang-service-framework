package server

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	Host   string
	Port   string
	Logger *logrus.Logger
}

func SetUpServer(host string, port string, logger *logrus.Logger) Server {
	server := Server{
		Host:   host,
		Port:   port,
		Logger: logger,
	}
	return server
}

type IServer interface {
	AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request))
	Start()
}

func (s Server) AddRoute(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	http.HandleFunc(path, handler)
}

func (s Server) Start() {
	s.Logger.Fatal(http.ListenAndServe(s.Host+":"+s.Port, nil))
}
