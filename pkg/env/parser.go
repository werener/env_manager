package env

import (
	"fmt"
	"strings"
)

// addError is a shorthand for attaching
// a new parsing error to the EnvFile.
func (file *EnvFile) addError(err error) {
	file.accumulatedErrors = append(file.accumulatedErrors, err)
}

// parse fills the specified EnvFile with Key-Value pairs
// and parsing errors.
func (file *EnvFile) parse() {
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
}

// parseExpr splits a provided expression into a Key-Value pair.
// It returns either a Key-Value pair or a parsing error.
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

// parseValue extracts a value from a provided string.
// It returns either the extracted value or a parsing error.
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
