// Package env provides and API for loading .env files.
package env

import (
	"github.com/werener/env_manager/internal/reader"
)

// EnvFile is a representation of the .env file.
// It contains its raw content, all parsing errors
// and Key-Value pairings that were extracted.
type EnvFile struct {
	env               map[string]string
	content           string
	accumulatedErrors []error
}

// Load creates a new representation of a .env file.
// It accepts a path to the .env file.
//
// A successful load returns an EnvFile instance
// and the first encountered non-parsing error, if any.
func Load(env_path string) (file EnvFile, err error) {
	content, openError := reader.ReadEnv(env_path)
	if openError != nil {
		return EnvFile{}, openError
	}

	file = newEnvFile(content)
	file.parse()

	return file, nil
}

// GetEnv returns all Key-Value pairs of the specified file.
func (file EnvFile) GetEnv() map[string]string {
	return file.env
}

// GetErrors returns all parsing errors of the specified file.
func (file EnvFile) GetErrors() []error {
	return file.accumulatedErrors
}

// newEnvFile is a shorthand to create a not yet
// parsed instance of a file from a provided content.
func newEnvFile(content string) EnvFile {
	return EnvFile{
		env:               map[string]string{},
		content:           content,
		accumulatedErrors: []error{},
	}
}
