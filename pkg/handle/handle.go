package handle

import (
	"flag"
	"fmt"
	"os"
)

type Command int

const (
	add Command = iota
	branch
	catFile
	commit
	help
)

func (c Command) String() string {
	switch c {
	case add:
		return "add"
	case branch:
		return "branch"
	case catFile:
		return "cat-file"
	case commit:
		return "commit"
	case help:
		return "help"
	default:
		return ""
	}
}

func parseArgs() {

	// Parse the command line arguments
	flag.Parse()

	command, args := flag.Arg(0), flag.Args()[1:]
	switch command {
	case add.String():
		handle(add.String(), args)
	case branch.String():
		handle(branch.String(), args)
	case catFile.String():
		handle(catFile.String(), args)
	case commit.String():
		handle(commit.String(), args)
	case help.String():
		handle(help.String(), args)
	default:
		fmt.Println("Unknown command: ", command)
		os.Exit(1)
	}
}

func handle(command string, args []string) {
	fmt.Println("Command: ", command)
	fmt.Println("Args: ", args)
}
