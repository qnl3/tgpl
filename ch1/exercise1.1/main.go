// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println("ex1.1time:", time.Since(start).Nanoseconds() )
}
