package main

import (
	"fmt"
	"strings"
	"testing"
)

// setUpWithTeardown
func setupWithTeardown() func() {
	fmt.Println("セットアップ")
	return func() {
		fmt.Println("Tear down")
	}
}

// TestToUpper はstringsのToUpperのテスト
func TestToUpper(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"foo", "FOO"},
		{"bar", "BAR"},
	}

	teardown := setupWithTeardown()
	t.Run("parent group", func(t *testing.T) {
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
	})
	teardown()
}

// TestToLower はstringsのToLowerのテスト
func TestToLower(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"FOO", "foo"},
		{"BAR", "bar"},
	}

	teardown := setupWithTeardown()
	t.Run("parent group", func(t *testing.T) {
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
	})
	teardown()
}
