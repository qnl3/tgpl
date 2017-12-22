package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type request struct {
	Method     string            `json:"method"`
	URL        string            `json:"url"`
	Protocol   string            `json:"protocol"`
	Headers    map[string]string `json:"headers"`
	Host       string            `json:"host"`
	RemoteAddr string            `json:"remoteAddress"`
	Form       map[string]string `json:"form,omitempty"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	count.increment()

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	//w.Header().Set("Content-Type", "apppliction/json")

	//log.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	h := make(map[string]string)
	for k, v := range r.Header {
		h[k] = strings.Join(v, "")
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	f := make(map[string]string)
	for k, v := range r.Form {
		f[k] = strings.Join(v, "")
	}

	req := request{
		Method:     r.Method,
		URL:        ("http://" + r.Host + r.URL.Path),
		Protocol:   r.Proto,
		Host:       r.Host,
		RemoteAddr: r.RemoteAddr,
		Headers:    h,
		Form:       f,
	}

	enc.Encode(req)
}
