package main

import (
	"log"
	"time"
)

func start(args []string) {
	start := time.Now()
	ch := make(chan string)
	for _, url := range args[1:] {
		go fetch(url, ch)
	}

	for range args[1:] {
		log.Println(<-ch)
	}

	log.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
