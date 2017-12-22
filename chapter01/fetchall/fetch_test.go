package main

import (
	"log"
	"testing"
)

func Test_fetch(t *testing.T) {
	type args struct {
		url string
		ch  chan string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "helloworld", args: args{url: "http://localhost:8080/helloworld", ch: make(chan string)}},
		{name: "cat", args: args{url: "http://localhost:8080/cat", ch: make(chan string)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go fetch(tt.args.url, tt.args.ch)
			log.Println(<-tt.args.ch)
		})
	}
}
