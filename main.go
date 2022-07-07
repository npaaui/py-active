package main

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"

	"active/app"
	. "active/common"
)

func main() {
	err := cleanenv.ReadConfig("./config/config.yml", &Cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(&Cfg)
}