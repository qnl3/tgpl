package main

import (
	"log"
	"net/http"
)

func serve(port int) {
	lissajousHandler := func(w http.ResponseWriter, r *http.Request) {
		incrementConter()
		lissajous(w)
	}
	http.HandleFunc("/", lissajousHandler)
	http.HandleFunc("/echo/", echoHandler)
	http.HandleFunc("/count", counterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
