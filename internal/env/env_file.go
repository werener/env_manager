package env

import (
	"github.com/werener/env_manager/internal/reader"
)

type EnvFile struct {
	env               map[string]string
	content           string
	AccumulatedErrors []error
}

func OpenEnvFile(env_path string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           reader.OpenEnv(env_path),
		AccumulatedErrors: []error{},
	}
}

func NewEnvFile(content string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           content,
		AccumulatedErrors: []error{},
	}
}

func (file *EnvFile) addError(err error) {
	file.AccumulatedErrors = append(file.AccumulatedErrors, err)
}
