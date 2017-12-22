package main

import "sync"

const port int = 8000

var (
	mu    sync.Mutex
	count int
)

func main() {
	serve(port)
}
