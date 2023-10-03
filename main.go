package main

import (
	"fmt"
	"telli/ArgParser"
	"telli/REPL"
)

// main: entry point of the program
func main() {
	ap := ArgParser.NewArgParser()
	if ap.REPL {
		// REPL
		fmt.Println("*** Starting Telli Interactive Mode ***")
		fmt.Println("Type 'exit()' to exit")
		REPL.R.RunREPL()
	} else {
		// compile file
		fmt.Println("compile file")
	}
}
