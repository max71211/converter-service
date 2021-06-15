package main

import (
	"log"

	"github.com/max71211/converter-service/pkg/config"
	"github.com/max71211/converter-service/pkg/server"

	"github.com/caarlos0/env/v6"
)

// @title YAML Converter Service
// @version 0.0.1
// @description This is the simple YAML Converter Service

// @contact.name Service Support
// @contact.email max71211@yandex.ru

// @host localhost:8080
// @BasePath /

// init is invoked before main()
func main() {
	if ok, err := run(); !ok {
		log.Fatal("Run error: ", err)
	}
}

func run() (ok bool, err error) {
	var cfg config.Config
	err = env.Parse(&cfg)
	if err != nil {
		return false, err
	}

	srv := server.NewServer(&cfg)

	err = srv.Run()
	if err != nil {
		return false, err
	}

	return true, nil
}
