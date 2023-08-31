package main

import (
	"avitoTask/internal/app"
	"avitoTask/internal/app/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	application := app.New(cfg)

	err = application.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)

}
