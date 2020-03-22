package utils

import (
	"encoding/json"
	"os"

	conf "github.com/atrariksa/golang-service-framework/src/models/config"
	c "github.com/atrariksa/golang-service-framework/src/constants"
)

var config conf.Config

func GetAppConfig(configDir string) conf.Config {
	file, err := os.Open(configDir)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(c.CONFIG_LOAD_ERROR)
	}
	return config
}
