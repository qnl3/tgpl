package main

import (
	"flag"
	"fmt"
	"strings"
)

func echo() {
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
