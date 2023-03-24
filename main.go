package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed usage/usage.txt
var usage string

func main() {
	input := os.Args[1:]

	// Check if no command is specified.
	if len(input) == 0 {
		fmt.Println("freeformgen: no command specified")
		fmt.Println(usage)
		return
	}

	switch cmd := input[0]; cmd {
	case "directives":
		exec, err := parseDirectiveCommand(input)
		if err != nil {
			fmt.Printf("freeformgen: %s\n", err)
			fmt.Println(directiveCommandUsage)
			return
		}
		err = exec.Handle()
		if err != nil {
			fmt.Println(err)
		}
	case "help":
		fmt.Println(usage)
	default:
		fmt.Printf("freeformgen: invalid command: `%s`\n", cmd)
		fmt.Println(usage)
	}

}
