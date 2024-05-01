package main

import (
	"terraform-roulette/cli"
	"terraform-roulette/roulette"
)

func main() {
	command, args := cli.ParseArgs()

	roulette.Play(command, args)
}
