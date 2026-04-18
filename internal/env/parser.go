package env

import (
	"fmt"
	"strings"
)

// Why no need for '*'??
func (file EnvFile) Parse() error {
	lineNumber := 1
	for expr := range strings.Lines(file.content) {
		lineNumber++

		if strings.HasPrefix(expr, "#") {
			continue
		}

		key, value, hasAssignment := strings.Cut(expr, "=")
		key, value = strings.TrimSpace(key), strings.TrimSpace(value)

		if len(key) == 0 {
			file.addError(fmt.Errorf("No key provided"))
		}
		if len(value) == 0 {
			file.addError(fmt.Errorf("No value provided"))
		}
		if !hasAssignment {
			file.addError(fmt.Errorf("Line without assignment"))
		}

		file.env[key] = value
	}
	return nil
}

func (file EnvFile) GetEnv() map[string]string {
	return file.env
}
