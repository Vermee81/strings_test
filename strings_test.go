package main

import (
	"strings"
	"testing"
)

// TestToUpper はstringsのToUpperのテスト
func TestToUpper(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"foo", "FOO"},
		{"bar", "BAR"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			result := strings.ToUpper(tt.input)
			if tt.expected != result {
				t.Errorf("expected %v, but %v", tt.expected, result)
			}
		})
	}
}

// TestToLower はstringsのToLowerのテスト
func TestToLower(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"FOO", "foo"},
		{"BAR", "bar"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			result := strings.ToLower(tt.input)
			if tt.expected != result {
				t.Errorf("expected %v, but %v", tt.expected, result)
			}
		})
	}
}
