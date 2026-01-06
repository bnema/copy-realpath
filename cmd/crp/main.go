package main

import (
	"fmt"
	"os"

	"github.com/bnema/copy-realpath/internal/app"
)

func main() {
	// Get input from args (default to empty string for current directory)
	var input string
	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	application := app.NewDefault()
	path, err := application.Run(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "crp: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(path)
}
