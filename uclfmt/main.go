package main

import (
	"fmt"
	"os"

	"github.com/gelraen/ucl"
)

func main() {
	v, err := ucl.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse input: %s\n", err)
		os.Exit(1)
	}
	err = ucl.Format(v, nil, os.Stdout)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to format input: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("\n")
}
