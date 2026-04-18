package reader

import (
	"os"
)

func OpenEnv(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
