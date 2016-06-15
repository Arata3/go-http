package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/facebookgo/grace/gracehttp"
)

func makeServe() {
	log.Printf("Start Port: http://127.0.0.1:%s\n", port)
	err := gracehttp.Serve(&http.Server{Addr: ":" + port, Handler: mux})
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	if err != nil {
		log.Fatal(err)
	}
}
