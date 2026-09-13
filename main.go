// Package main is the entry point for the githubaction demo service.
//
// This file exists to give the CI pipeline (.github/workflows/ci.yml)
// something real to lint, test, and build against -- a pipeline with
// nothing behind it can't prove anything.
package main

import (
	"fmt"
	"os"
)

// add returns the sum of two integers.
//
// Kept trivial and pure (no I/O, no globals) so it's a clean example for
// go test -race -cover: deterministic, no shared state to race on.
func add(a, b int) int {
	return a + b
}

func main() {
	result := add(2, 3)

	// fmt.Fprintln to stdout rather than fmt.Println is a deliberate habit:
	// it makes the output stream explicit, which matters the moment this
	// grows past one print statement and some output needs to go to stderr
	// instead (errors, diagnostics) without a silent behavior change.
	if _, err := fmt.Fprintln(os.Stdout, "2 + 3 =", result); err != nil {
		// Failing to write to stdout is rare but not impossible (closed
		// pipe, full disk on the receiving end). Exit non-zero rather than
		// silently swallow it -- a CI job piping this output should see
		// the failure, not a false green.
		fmt.Fprintln(os.Stderr, "writing output:", err)
		os.Exit(1)
	}
}
