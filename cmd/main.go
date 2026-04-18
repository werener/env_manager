package main

import (
	"fmt"
	"os"

	"github.com/werener/env_manager/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Expected a path to env file to be provided\nUsage: %s [ENV]\n", os.Args[0])
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Printf("Expected only one path to env file to be provided\nUsage: %s [ENV]\n", os.Args[0])
		os.Exit(1)
	}

	if err := run(os.Args[1]); err != nil {
		fmt.Println()
	}
}

func run(envPath string) error {
	file := parser.NewEnvFile(envPath)
	err := file.Parse()

	return err
}
