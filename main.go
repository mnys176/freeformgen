package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/mnys176/usage"
)

type freeformgenError struct {
	err error
}

func (e freeformgenError) Error() string {
	return "freeformgen: " + e.err.Error()
}

func init() {
	usage.Init("freeformgen")
	usage.AddArg("<path-to-source>")

	outputOption, _ := usage.NewOption([]string{"--output", "-o"}, "Desired output path for the generated directives with respect to the current working directory. By default, \"./directives_gen.go\" is used.")
	outputOption.AddArg("<path>")
	usage.AddOption(outputOption)

	helpOption, _ := usage.NewOption([]string{"--help", "-h"}, "Prints this help message and exits.")
	usage.AddOption(helpOption)

	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usage.Usage())
	}
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		err := freeformgenError{errors.New("no source file or directory provided")}
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	// // Check if no command is specified.
	// if len(flag.Args()) == 0 {
	// 	fmt.Fprintln(os.Stderr, globals.NoCommandError())
	// 	flag.Usage()
	// 	return
	// }

	// switch cmd := flag.Arg(0); cmd {
	// case globals.DirectivesCommand:
	// 	err := directives.Parse()
	// 	if err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		return
	// 	}
	// 	if err = directives.Execute(); err != nil {
	// 		fmt.Fprintln(os.Stderr, err)
	// 		return
	// 	}
	// case globals.HelpCommand:
	// 	flag.Usage()
	// default:
	// 	fmt.Fprintln(os.Stderr, globals.UnknownCommandError(cmd))
	// }
}
