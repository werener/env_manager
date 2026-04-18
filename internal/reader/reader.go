package reader

import (
	"os"
)

func ReadEnv(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
