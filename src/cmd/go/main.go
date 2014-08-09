package main

import "dummy0"

func main() {
	// Test that we can refer to any pkg defined in the stdlib repo, not just
	// pkgs from a predefined list. This tests that if the stdlib adds pkgs or
	// definitions, we'll analyze them instead of just analyzing the code in
	// GOROOT.

	v := dummy0.MyFunc()
	_ = v
}
