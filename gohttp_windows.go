package main

import (
	"log"
	"net/http"
)

func makeServe() {
	log.Printf("Start Port: http://127.0.0.1:%s\n", port)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
