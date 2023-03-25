package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/mnys176/freeformgen/directives"
)

type executable interface {
	Execute() error
}

type freeformgenError struct {
	Message string
}

func (e freeformgenError) Error() string {
	return fmt.Sprintf("freeformgen: %s", e.Message)
}

//go:embed usage.txt
var usage string

func execute(e executable) {
	if err := e.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func noCommandSpecifiedError() error {
	return freeformgenError{"no command specified"}
}

func unknownCommandError(command string) error {
	return freeformgenError{fmt.Sprintf("unknown command: %s", command)}
}

func init() {
	flag.Usage = func() {
		// TODO: Dynamically generate this.
		fmt.Fprintln(os.Stdout, usage)
	}
}

func main() {
	flag.Parse()

	// Check if no command is specified.
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, noCommandSpecifiedError())
		flag.Usage()
		return
	}

	var action executable
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "directives":
		action, err = directives.Parse()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case "help":
		flag.Usage()
	default:
		fmt.Fprintln(os.Stderr, unknownCommandError(cmd))
	}
	execute(action)
}
