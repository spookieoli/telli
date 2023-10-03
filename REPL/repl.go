package REPL

import "fmt"

// REPL: struct to handle REPL
type REPL struct {
}

// RunREPL: runs the REPL
func (r *REPL) RunREPL() {
	prompt := ":>>> "
	command := ""
	for command != "exit()" {
		fmt.Print(prompt)
		// read in next command from console
		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Println("Error reading command")
		} else {
			// TODO: lex and parse command
			fmt.Println("Command: " + command)
		}
	}
}
