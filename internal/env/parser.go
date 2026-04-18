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

		key, value, parseErr := parseExpr(expr)
		if parseErr != nil {
			file.addError(fmt.Errorf("%s: Line %d", parseErr, lineNumber))
			continue
		}

		file.env[key] = value
	}

	return nil
}

func parseExpr(expr string) (string, string, error) {
	key, value, hasAssignment := strings.Cut(expr, "=")
	if !hasAssignment {
		return "", "", fmt.Errorf("Non-empty line without assignment")
	}

	key = strings.TrimSpace(key)
	value, valueParseErr := parseValue(value)

	if valueParseErr != nil {
		return "", "", valueParseErr
	}
	if len(key) == 0 {
		return "", "", fmt.Errorf("No key provided")
	}
	if len(value) == 0 {
		return "", "", fmt.Errorf("No value provided")
	}

	return key, value, nil
}

func parseValue(value string) (string, error) {
	value = strings.TrimSpace(value)
	if strings.HasPrefix(value, "\"") {
		value = strings.TrimPrefix(value, "\"")
		if !strings.HasSuffix(value, "\"") {
			return "", fmt.Errorf("Unterminated quote")
		}
		value = strings.TrimSuffix(value, "\"")
	}
	return value, nil
}

func (file EnvFile) GetEnv() map[string]string {
	return file.env
}
