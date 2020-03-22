package main

import (
	"os"
	"path/filepath"
	"strconv"

	c "github.com/atrariksa/golang-service-framework/src/constants"
	log "github.com/atrariksa/golang-service-framework/src/logger"
	r "github.com/atrariksa/golang-service-framework/src/routers"
	s "github.com/atrariksa/golang-service-framework/src/server"
	"github.com/atrariksa/golang-service-framework/src/utils"
)

func main() {

	// get current running directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(c.UNEXPECTED_ERROR)
	}

	// initialize log and load config
	logger := log.GetLogger(dir + c.LOG_FILE_PATH)
	config := utils.GetAppConfig(dir + c.CONFIG_FILE_PATH)

	// set up server
	host := config.AppConfig.HttpServerConfig.Host
	port := strconv.FormatInt(config.AppConfig.HttpServerConfig.Port, 10)
	server := s.SetUpServer(host, port, logger)

	// setup route
	router := r.SetUpRouter(server, logger)
	router.Route()

	// set field server info
	fields := make(map[string]interface{})
	fields["host_field"] = host
	fields["port_field"] = port

	// starting server
	logger.WithFields(fields).Info(c.START_UP)
	server.Start()
}
