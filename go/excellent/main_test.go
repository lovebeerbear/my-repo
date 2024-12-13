package main

import (
	"bytes"
	"testing"
)

func TestMain(t *testing.T) {
	var buf bytes.Buffer
	version = "1.0.0"
	expectedOutput := "Version:1.0.0\n"

	// Redirect stdout
	old := fmt.Println
	defer func() { fmt.Println = old }()
	fmt.Println = func(a ...interface{}) (n int, err error) {
		return buf.WriteString(fmt.Sprintln(a...))
	}

	main()

	if buf.String() != expectedOutput {
		t.Errorf("expected %q but got %q", expectedOutput, buf.String())
	}
}