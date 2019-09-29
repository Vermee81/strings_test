package main

import (
	"strings"
	"testing"
)

func Test_ToUpper(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"foo", "FOO"},
		{"bar", "BAR"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			result := strings.ToUpper(tt.input)
			if tt.expected != result {
				t.Errorf("expected %v, but %v", tt.expected, result)
			}
		})
	}
}
