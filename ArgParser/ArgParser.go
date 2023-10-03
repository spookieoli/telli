package ArgParser

import "os"

// ArgParser: struct to parse command line arguments
type ArgParser struct {
	args   []string
	length int
	REPL   bool
}

// NewArgParser: constructor for ArgParser
func NewArgParser() *ArgParser {
	args := os.Args
	if len(args) == 2 {
		if _, err := os.Stat(args[1]); err == nil {
			// if file exists and is a path
			return &ArgParser{args, len(args), false}
		}
	}
	// else REPL
	return &ArgParser{args, len(args), true}
}
