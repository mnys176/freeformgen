package directives

import (
	"flag"

	"github.com/mnys176/freeformgen/globals"
)

func Parse() error {
	fs := flag.NewFlagSet(globals.DirectivesCommand, flag.ContinueOnError)

	fs.BoolVar(&directory, "directory", false, "")
	fs.BoolVar(&directory, "d", false, "")
	fs.BoolVar(&help, "help", false, "")
	fs.BoolVar(&help, "h", false, "")

	// Parse options.
	fs.Parse(flag.Args()[1:])

	// Parse arguments.
	if n := fs.NArg(); n > 1 {
		return globals.IncorrectNumberOfArgumentsError(n)
	}
	packageSourcePath = fs.Arg(0)
	return nil
}
