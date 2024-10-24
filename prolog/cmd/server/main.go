package main

import (
	"log"

	// "github.com/Avinash-Vadlamudi/proglog/internal/server"
	"github.com/Avinash-Vadlamudi/prolog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
