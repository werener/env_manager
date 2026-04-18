package parser

import "github.com/werener/env_manager/internal/reader"

type EnvFile struct {
	env               map[string]string
	content           string
	accumulatedErrors []error
}

func NewEnvFile(env_path string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           reader.OpenEnv(env_path),
		accumulatedErrors: []error{},
	}
}
