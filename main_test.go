package main

import "testing"

// Table-driven so adding a new case is a one-line addition, not a new
// function -- the pattern the pipeline's `go test ./...` is built to run
// as the codebase grows past this single example.
func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		a, b     int
		expected int
	}{
		{"positive numbers", 2, 3, 5},
		{"negative numbers", -2, -3, -5},
		{"zero", 0, 0, 0},
		{"mixed signs", -5, 10, 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := add(tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("add(%d, %d) = %d, want %d", tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
