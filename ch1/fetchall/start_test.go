package main

import "testing"

func Test_start(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{name: "helloworld", args: []string{"", "http://localhost:8080/helloworld", "http://localhost:8080/cat"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start(tt.args)
		})
	}
}
