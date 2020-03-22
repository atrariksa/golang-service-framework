package config

type HttpServerConfig struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}
