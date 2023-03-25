package directives

import (
	"flag"
)

func Parse() (*DirectivesCommand, error) {
	// Create executable.
	dc := DirectivesCommand{}
	DirectivesFlagSet.BoolVar(&dc.Directory, "directory", false, "make me better")
	DirectivesFlagSet.BoolVar(&dc.Directory, "d", false, "make me better")

	// Parse flags.
	DirectivesFlagSet.Parse(flag.Args()[1:])

	// Parse arguments.
	if n := DirectivesFlagSet.NArg(); n > 1 {
		return nil, incorrectNumberOfArgumentsError(n)
	}
	dc.Path = DirectivesFlagSet.Arg(0)
	return &dc, nil
}
