package main

import (
	"fmt"
	"os"

	"github.com/werener/env_manager/pkg/env"
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
	file, err := env.Load(envPath)

	fmt.Printf("Erorrs, that occured during parsing %s:\n", envPath)
	for _, e := range file.GetErrors() {
		fmt.Println(e)
	}
	fmt.Printf("\nKey-Value pairs, that were extracted from %s:\n", envPath)
	for k, v := range file.GetEnv() {
		fmt.Printf("%s - %s\n", k, v)
	}

	return err
}
