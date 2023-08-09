package main

import (
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

func (e freeformgenError) Is(target error) bool {
	return e.Error() == target.Error()
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

	// for i := 0; i < 1000; i++ {
	// 	fmt.Println(vectorDirective("", 1, 3))
	// }
	// args := flag.Args()

	// if len(args) == 0 {
	// 	err := freeformgenError{errors.New("no source file or directory provided")}
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// 	return
	// }

	// for _, arg := range args {
	// 	fmt.Printf("%q\n", arg)
	// }

	fmt.Println("------------------------------")

}
