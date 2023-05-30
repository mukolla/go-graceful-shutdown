package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HttpAddr string `envconfig:"MYAPP_HTTP_ADDR" default:"127.0.0.1:8000"`
	DBAddr   string `envconfig:"MYAPP_DB_ADDR"`
}

func NewConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("myapp", &c)
	if err != nil {
		log.Fatal(err.Error())
	}

	//format := "Debug: \nHttpAddr: %s\nDBAddr: %s\n"
	//_, err = fmt.Printf(format, c.HttpAddr, c.DBAddr)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
