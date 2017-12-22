// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for i, arg := range os.Args[1:] {
		fmt.Printf("%v: %s\n", i, arg)
	}
	fmt.Println("ex1.2 time",time.Since(start).Nanoseconds())
}