package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type counter struct {
	Count int `json:"count"`
}

var (
	mu    sync.Mutex
	count = counter{Count: 0}
)

func counterHandler(w http.ResponseWriter, r *http.Request) {

	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	//w.Header().Set("Content-Type", "apppliction/json")

	mu.Lock()
	enc.Encode(count)
	mu.Unlock()
}

func (c counter) increment() {
	mu.Lock()
	count.Count++
	mu.Unlock()
}
