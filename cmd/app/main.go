package main

import (
	"log"

	"github.com/koizumi55555/corporation-api/config"
	"github.com/koizumi55555/corporation-api/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.Run(cfg); err != nil {
		log.Fatalf("runtime error %s", err)
	}
}
