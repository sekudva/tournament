package main

import (
	"fmt"
	"os"

	"github.com/sekudva/strategika/cmd/cli"
)

func main() {
	cli.ParseFlags()

	if err := cli.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
