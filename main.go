package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/mnys176/freeformgen/directives"
	"github.com/mnys176/freeformgen/globals"
)

//go:embed usage.txt
var usage string

func usageFunc() {
	fmt.Fprintln(os.Stdout, usage)
}

func main() {
	flag.Usage = usageFunc
	flag.Parse()

	// Check if no command is specified.
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, globals.NoCommandError())
		flag.Usage()
		return
	}

	switch cmd := flag.Arg(0); cmd {
	case globals.DirectivesCommand:
		err := directives.Parse()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		if err = directives.Execute(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	case globals.HelpCommand:
		flag.Usage()
	default:
		fmt.Fprintln(os.Stderr, globals.UnknownCommandError(cmd))
	}
}
