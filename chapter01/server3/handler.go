package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		log.Printf("Header[%q] = %q\n", k, v)
	}
	log.Printf("Host = %q\n", r.Host)
	log.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		log.Printf("Form[%q] = %q\n", k, v)
	}

}
