package main

import (
	"fmt"
	"log"

	"github.com/<TEMPLATE>/config"
	"github.com/<TEMPLATE>/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Start(cfg)

	fmt.Print("Hello Wold!!")
}
