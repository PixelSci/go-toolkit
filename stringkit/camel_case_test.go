package stringkit

import "testing"

func TestCamelCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "helloWorld"},
		{"some-hyphen-text", "someHyphenText"},
		{"CONSTANT_CASE", "constantCase"},
		{"PascalCase", "pascalCase"},
		{"mixed   SpAcE", "mixedSpAcE"},
	}

	for _, tt := range tests {
		got := CamelCase(tt.input)
		if got != tt.expected {
			t.Errorf("CamelCase(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}
