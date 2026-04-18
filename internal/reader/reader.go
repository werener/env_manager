package reader

import (
	"log"
	"os"
)

func OpenEnv(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
