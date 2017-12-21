package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	log.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	log.Fprintf(s, "%s %s %s\n", r.Method, r.Url, r.Proto)

	for k, v := range r.Header {
		log.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
