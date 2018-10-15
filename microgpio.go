package main

import (
	"log"

	"github.com/teran/microgpio/server"
)

// Version var to be set by linker
var Version = "dev build"

func main() {
	log.Printf("Starting microgpio=%s", Version)

	srv := server.New()
	log.Fatalf("%s", srv.ListenAndServe(":8080"))
}
