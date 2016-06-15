package main

import (
	"log"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
)

func makeServe() {
	log.Printf("Start Port: http://127.0.0.1:%s\n", port)
	err := gracehttp.Serve(&http.Server{Addr: ":" + port, Handler: mux})
	if err != nil {
		log.Fatal(err)
	}
}
