package reader

import (
	"log"
	"os"
	"strings"
)

func OpenEnv(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(data))
}
