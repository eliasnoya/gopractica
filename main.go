package main

import (
	"core/config"
	"core/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

type HomeStruct struct {
	PageTitle string
}

// Start HTTP Server
func main() {

	fmt.Println("================================================")
	fmt.Println("Starting server on PORT: " + config.Get().HttpPort)
	fmt.Println("================================================")

	srv := &http.Server{
		Handler: router.GetRouter(),
		Addr:    "127.0.0.1:" + config.Get().HttpPort,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
