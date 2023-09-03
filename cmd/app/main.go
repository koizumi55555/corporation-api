package main

import (
	"koizumi55555/go-restapi/config"
	"koizumi55555/go-restapi/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("")
	}
	if err := app.Run(cfg); err != nil {
		log.Fatalf("runtime error %s", err)
	}
}
