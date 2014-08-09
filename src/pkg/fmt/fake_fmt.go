package fmt

// FakePrintf is not the real printf. We're testing that our analyzer looks at
// the repository being analyzed, and not the Go stdlib, if GoBaseImportPaths is
// set.
func FakePrintf(s string) {}
