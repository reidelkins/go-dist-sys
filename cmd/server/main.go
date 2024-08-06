package main

import (
	"log"

	"github.com/reidelkins/go-dist-sys/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}