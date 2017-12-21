package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mu    sync.Mutex
	count int
)

func counterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

func incrementConter() {
	mu.Lock()
	count++
	mu.Unlock()
}
