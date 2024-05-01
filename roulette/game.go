package roulette

import (
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var options = map[string]int{
	"destroy":              25,
	"apply --auto-approve": 50,
}

func assertTerraformCommand(command string) {
	cmd := exec.Command("terraform", command, "--help")
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func execTerraform(commands []string, args []string) {
	cmdArgs := append(commands, args...)
	cmd := exec.Command("terraform", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}
}

func spinTheWheel(commands []string) []string {
	for option, weight := range options {
		if rand.Intn(100) < weight {
			commands = strings.Split(option, " ")
			break
		}
	}

	return commands
}

func Play(command string, args []string) {
	assertTerraformCommand(command)

	commands := strings.Split(command, " ")

	if !strings.HasPrefix(command, "apply") {
		execTerraform(commands, args)
		os.Exit(0)
	}

	commands = spinTheWheel(commands)

	execTerraform(commands, args)
}
