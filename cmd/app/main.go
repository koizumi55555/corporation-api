package main

import (
	"koizumi55555/corporation-api/config"
	"koizumi55555/corporation-api/internal/app"
	"log"
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
