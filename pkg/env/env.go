package env

import (
	"github.com/werener/env_manager/internal/reader"
)

type EnvFile struct {
	env               map[string]string
	content           string
	accumulatedErrors []error
}

func Load(env_path string) (EnvFile, error) {
	content, openError := reader.ReadEnv(env_path)
	if openError != nil {
		return EnvFile{}, openError
	}

	file := newEnvFile(content)
	parseError := file.parse()

	return file, parseError
}

func (file EnvFile) GetEnv() map[string]string {
	return file.env
}
func (file EnvFile) GetErrors() []error {
	return file.accumulatedErrors
}

func newEnvFile(content string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           content,
		accumulatedErrors: []error{},
	}
}
