package main

import (
	"log"
	"net/http"
)

func serve(port int) {
	http.HandleFunc("/", lissajousHandler)
	http.HandleFunc("/echo/", echoHandler)
	http.HandleFunc("/count", counterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
