package directives

import (
	_ "embed"
	"flag"
	"fmt"
	"os"

	"github.com/mnys176/freeformgen/globals"
)

//go:embed usage.txt
var usage string

func usageFunc() {
	fmt.Fprintln(os.Stdout, usage)
}

func Parse() error {
	fs := flag.NewFlagSet(globals.DirectivesCommand, flag.ContinueOnError)
	fs.Usage = usageFunc

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
