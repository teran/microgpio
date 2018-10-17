package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/teran/microgpio/controller"
	"github.com/teran/microgpio/server"
)

// Version var to be set by linker
var Version = "dev build"

type config struct {
	ListenAddr string         `envconfig:"LISTEN_ADDR" default:":8080"`
	PinMapping map[string]int `envconfig:"PIN_MAPPING" required:"true"`
}

func main() {
	var cfg config
	envconfig.MustProcess("microgpio", &cfg)

	log.Printf("Starting microgpio=%s", Version)

	c, err := controller.New(cfg.PinMapping)
	if err != nil {
		log.Fatalf("error initializing controller: %s", err)
	}

	srv := server.New(c)
	log.Fatalf("%s", srv.ListenAndServe(cfg.ListenAddr))
}
