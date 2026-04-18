package parser

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func (file EnvFile) Parse() {
	scanner := bufio.NewScanner(strings.NewReader(file.content))
	for scanner.Scan() {
		expr := scanner.Text()

		if strings.HasPrefix(expr, "#") {
			continue
		}

		key, value, hasAssignment := strings.Cut(expr, "=")
		if !hasAssignment {
			continue
		}

		file.variables[key] = value

		fmt.Println(expr, "+")
	}

	for k, v := range file.variables {
		fmt.Printf("Key: %s\nValue: %s\n\n", k, v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
