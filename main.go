package main

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/status", HandlerStatus)

	s := &http.Server{
		Addr:           ":4443",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http2.ConfigureServer(s, &http2.Server{})

	log.Fatal(s.ListenAndServeTLS("server.crt", "server.key"))
}
