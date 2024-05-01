package cli

import (
	"os"
)

func ParseArgs() (string, []string) {
	command := ""
	args := []string{}

	if len(os.Args) > 1 {
		command = os.Args[1]
		args = os.Args[2:]
	}

	return command, args
}
