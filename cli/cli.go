package cli

import (
	"os"
	"terraform-roulette/roulette"
)

func RunCli() {
	command := os.Args[1]
	args := os.Args[2:]

	roulette.Play(command, args)
}
