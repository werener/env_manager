package env

import (
	"github.com/werener/env_manager/internal/reader"
)

type EnvFile struct {
	env               map[string]string
	content           string
	accumulatedErrors []error
}

func OpenEnvFile(env_path string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           reader.OpenEnv(env_path),
		accumulatedErrors: []error{},
	}
}

func NewEnvFile(content string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           content,
		accumulatedErrors: []error{},
	}
}

func (file *EnvFile) addError(err error) {
	file.accumulatedErrors = append(file.accumulatedErrors, err)
}
