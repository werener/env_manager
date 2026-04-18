package main

import (
	"env_manager/internal/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Expected a path to env file to be provided\nUsage: %s [ENV]\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Printf("Expected only one path to env file to be provided\nUsage: %s [ENV]\n", os.Args[0])
		os.Exit(1)
	}

	run(os.Args[1])
}

func run(envPath string) {
	file := parser.NewEnvFile(envPath)
	file.Parse()
}
