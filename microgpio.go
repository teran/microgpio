package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"

	"github.com/teran/microgpio/server"
)

// Version var to be set by linker
var Version = "dev build"

type config struct {
	ListenAddr string `envconfig:"LISTEN_ADDR" default:":8080"`
}

func main() {
	var cfg config
	envconfig.MustProcess("microgpio", &cfg)

	log.Printf("Starting microgpio=%s", Version)

	srv := server.New()
	log.Fatalf("%s", srv.ListenAndServe(cfg.ListenAddr))
}
