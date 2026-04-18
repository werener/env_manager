package env

import (
	"fmt"
	"strings"
)

func (file *EnvFile) Parse() error {
	lineNumber := 0
	for expr := range strings.Lines(file.content) {
		lineNumber++
		expr := strings.TrimSpace(expr)

		if len(expr) == 0 || strings.HasPrefix(expr, "#") {
			continue
		}

		key, value, hasAssignment := strings.Cut(expr, "=")
		if !hasAssignment {
			file.addError(fmt.Errorf("Line without assignment. Line: %d", lineNumber))
			continue
		}

		key = strings.TrimSpace(key)
		value, valueErr := parseValue(value, lineNumber)
		if valueErr != nil {
			file.addError(valueErr)
			continue
		}

		if len(key) == 0 {
			file.addError(fmt.Errorf("No key provided. Line: %d", lineNumber))
			continue
		}

		if len(value) == 0 {
			file.addError(fmt.Errorf("No value provided. Line: %d", lineNumber))
			continue
		}

		file.env[key] = value
	}

	return nil
}

func parseValue(value string, lineNumber int) (string, error) {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "\"") {
		value = strings.TrimPrefix(value, "\"")
		if !strings.HasSuffix(value, "\"") {
			return "", fmt.Errorf("Unterminated quote: %d", lineNumber)
		}
		value = strings.TrimSuffix(value, "\"")
	}
	return value, nil
}

func (file EnvFile) GetEnv() map[string]string {
	return file.env
}
