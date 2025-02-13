package gitignore

import (
	"testing"
)

func TestProcessLines(t *testing.T) {
	tests := []struct {
		name         string
		keepComments bool
		keepNewLines bool
		input        []string
		expected     []string
	}{
		{
			name:         "Keep comments and newlines",
			keepComments: true,
			keepNewLines: true,
			input:        []string{"# comment", "node_modules", "", "node_modules", "dist"},
			expected:     []string{"# comment", "node_modules", "", "dist"},
		},
		{
			name:         "Remove comments and newlines",
			keepComments: false,
			keepNewLines: false,
			input:        []string{"# comment", "node_modules", "", "node_modules", "dist"},
			expected:     []string{"node_modules", "dist"},
		},
		{
			name:         "Remove duplicates",
			keepComments: true,
			keepNewLines: true,
			input:        []string{"node_modules", "node_modules", "dist", "dist"},
			expected:     []string{"node_modules", "dist"},
		},
		{
			name:         "Empty input",
			keepComments: true,
			keepNewLines: true,
			input:        []string{},
			expected:     []string{},
		},
		{
			name:         "Only comments",
			keepComments: true,
			keepNewLines: true,
			input:        []string{"# comment1", "# comment2"},
			expected:     []string{"# comment1", "# comment2"},
		},
		{
			name:         "Only empty lines",
			keepComments: true,
			keepNewLines: true,
			input:        []string{"", "", ""},
			expected:     []string{"", "", ""},
		},
		{
			name:         "Mixed content with duplicates",
			keepComments: true,
			keepNewLines: false,
			input:        []string{"# comment", "node_modules", "", "node_modules", "dist"},
			expected:     []string{"# comment", "node_modules", "dist"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Processor{
				KeepComments: tt.keepComments,
				KeepNewLines: tt.keepNewLines,
			}

			result := p.ProcessLines(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("expected %d lines, got %d", len(tt.expected), len(result))
			}

			for i, line := range result {
				if line != tt.expected[i] {
					t.Errorf("line %d: expected %q, got %q", i, tt.expected[i], line)
				}
			}
		})
	}
}
