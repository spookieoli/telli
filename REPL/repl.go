package REPL

import (
	"fmt"
	"telli/Tokenizer"
)

// REPL: struct to handle REPL
type REPL struct {
	Tokenizer Tokenizer.Tokenizer
}

var R REPL

func init() {
	R = REPL{Tokenizer: Tokenizer.T}
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
			fmt.Println(r.Tokenizer.Tokenize(command))
			// TODO: output result
		}
	}
}
