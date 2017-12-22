// Echo1 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	n   = flag.Bool("n", false, "Omit trailing newline.")
	sep = flag.String("s", " ", "seperator")
)

func main() {
	start := time.Now()
	flag.Parse()
	echo()
	fmt.Print("ec4 time:", time.Since(start).Nanoseconds())
}
