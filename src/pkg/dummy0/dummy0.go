package dummy0

import "dummy1"

// Refs to dummy1.MyType test that we can refer to any pkg defined in this repo
// as a stdlib package, not just pkgs from a predefined list. This tests that if
// the stdlib adds pkgs or definitions, we'll analyze them instead of just
// analyzing the code in GOROOT.

func MyFunc() dummy1.MyType {
	return dummy1.MyType(123)
}
