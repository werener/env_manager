package parser

import "github.com/werener/env_manager/internal/reader"

type EnvFile struct {
	variables map[string]string
	content   string
}

func NewEnvFile(env_path string) EnvFile {
	return EnvFile{
		content:   reader.OpenEnv(env_path),
		variables: map[string]string{},
	}
}
