package parser

import (
	"fmt"
	"strings"
)

func (file EnvFile) Parse() error {
	for expr := range strings.Lines(file.content) {

		if strings.HasPrefix(expr, "#") {
			continue
		}

		key, value, hasAssignment := strings.Cut(expr, "=")
		key, value = strings.TrimSpace(key), strings.TrimSpace(value)

		if len(key) == 0 {
			file.accumulatedErrors = append(file.accumulatedErrors, fmt.Errorf(""))
		}

		if !hasAssignment {
			continue
		}

		file.env[key] = value
	}
	return nil
}
