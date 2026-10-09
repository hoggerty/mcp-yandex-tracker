package cmd

import (
	"fmt"
	"os"
)

// errorf prints a prefixed error message to stderr.
func errorf(format string, args ...any) {
	_, _ = fmt.Fprintf(os.Stderr, "yandex-tracker-mcp: "+format+"\n", args...)
}

// fatalf prints a prefixed error message to stderr and exits with status 1.
func fatalf(format string, args ...any) {
	errorf(format, args...)
	os.Exit(1)
}
