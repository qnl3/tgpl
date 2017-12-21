package main

import (
	"log"
	"net/http"
)

func serve(port int) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
