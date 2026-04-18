package env

import (
	"testing"
)

func TestEnvFile_Parse(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expectedEnv map[string]string
	}{
		{
			name:    "basic key-value pairs",
			content: "KEY=val\nVAL=key",
			expectedEnv: map[string]string{
				"KEY": "val",
				"VAL": "key",
			},
		},
		{
			name:    "with comments",
			content: "KEY=val\nVAL=key",
			expectedEnv: map[string]string{
				"KEY": "val",
				"VAL": "key",
			},
		},
		{
			name:        "empty file",
			content:     "",
			expectedEnv: map[string]string{},
		},
		{
			name:    "value contains equals sign",
			content: "URL=https://example.com?foo=bar",
			expectedEnv: map[string]string{
				"URL": "https://example.com?foo=bar",
			},
		},
		{
			name:    "file has empty lines",
			content: "\nHELLO=WORLD\n    \t\nGOODBYE=WORLD!   ",
			expectedEnv: map[string]string{
				"HELLO":   "WORLD",
				"GOODBYE": "WORLD!",
			},
		},
		{
			name:    "obscolete spaces inside expression",
			content: "\nHELLO = WORLD\nGOODBYE =  WORLD!   ",
			expectedEnv: map[string]string{
				"HELLO":   "WORLD",
				"GOODBYE": "WORLD!",
			},
		},
		{
			name:    "with non-assigning lines",
			content: "FOO=bar\nINVALID_LINE\nBAZ=qux",
			expectedEnv: map[string]string{
				"FOO": "bar",
				"BAZ": "qux",
			},
		},
		{
			name:    "with quotation marks",
			content: "\nHELLO=\"WORLD\"\n    \t\nGOODBYE=\"WORLD!\"",
			expectedEnv: map[string]string{
				"HELLO":   "WORLD",
				"GOODBYE": "WORLD!",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := NewEnvFile(tt.content)
			file.Parse()

			if len(file.env) != len(tt.expectedEnv) {
				t.Errorf("got %d variables, want %d", len(file.env), len(tt.expectedEnv))
			}

			for key, expectedVal := range tt.expectedEnv {
				actualVal, ok := file.env[key]

				if !ok {
					t.Errorf("missing key %q", key)
					continue
				}
				if actualVal != expectedVal {
					t.Errorf("\nkey: %q got: %q, expected: %q", key, actualVal, expectedVal)
				}
			}
		})
	}
}
