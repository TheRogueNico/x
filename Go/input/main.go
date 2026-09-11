package main

import (
	"fmt"
	"io"
	"os"
)

func process(r io.Reader) error {
	// TODO: Replace with custom logic.
	_, err := io.Copy(os.Stdout, r)
	return err
}

func run(args []string) error {
	// No args: read stdin.
	if len(args) == 0 {
		return process(os.Stdin)
	}

	var firstErr error
	for _, name := range args {
		if err := processArg(name); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", name, err)
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	return firstErr
}

func processArg(name string) error {
	if name == "-" {
		return process(os.Stdin)
	}

	f, err := os.Open(name)
	if err != nil {
		return err
	}
	defer f.Close()

	return process(f)
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
