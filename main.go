package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"text/template"

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

	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"foo": func(op string, args ...any) (any, error) {
			switch op {
			case "one":
				if len(args) != 1 {
					return nil, errors.New("wrong number of arguments")
				}
				return args[0].(int), nil
			case "two":
				if len(args) != 2 {
					return nil, errors.New("wrong number of arguments")
				}
				return args[0].(int) + args[1].(int), nil
			case "three":
				if len(args) != 3 {
					return nil, errors.New("wrong number of arguments")
				}
				return args[0].(int) + args[1].(int) + args[2].(int), nil
			default:
				return nil, errors.New("invalid operation")
			}
		},
	}).Parse(`{{ foo "two" 1 2}}` + "\n"))

	err := tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}

	// args := flag.Args()

	// if len(args) == 0 {
	// 	err := freeformgenError{errors.New("no source file or directory provided")}
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// 	return
	// }

	// for _, arg := range args {
	// 	fmt.Printf("%q\n", arg)
	// }

	// fmt.Println("------------------------------")

}
