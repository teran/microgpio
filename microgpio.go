package main

import (
	"log"

	"github.com/teran/microgpio/drivers/fake"
	"github.com/teran/microgpio/server"
)

func main() {
	drv := &fake.FakeDriver{
		HighFunc: func(id int) error {
			log.Printf("Set High to %d", id)
			return nil
		},
		LowFunc: func(id int) error {
			log.Printf("Set Low to %d", id)
			return nil
		},
		OutputFunc: func(id int) error {
			log.Printf("Set Output to %d", id)
			return nil
		},
	}

	srv := server.New(drv)
	log.Fatalf("%s", srv.ListenAndServe("127.0.0.1:8080"))
}
