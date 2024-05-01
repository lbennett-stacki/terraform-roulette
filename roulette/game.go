package roulette

import (
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

var overrides = map[string]int{
	"destroy":              25,
	"apply --auto-approve": 50,
}

var overrideables = []string{"apply", "plan", "destroy"}

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
	target := rand.Intn(100)

	for override, weight := range overrides {
		if target < weight {
			commands = strings.Split(override, " ")
			break
		}
	}

	return commands
}

func isOverridable(command string) bool {
	result := false

	for _, overrideable := range overrideables {
		if command == overrideable {
			result = true
			break
		}
	}

	return result
}

func Play(command string, args []string) {
	assertTerraformCommand(command)

	commands := strings.Split(command, " ")

	if !isOverridable(command) {
		execTerraform(commands, args)
		os.Exit(0)
	}

	commands = spinTheWheel(commands)

	execTerraform(commands, args)
}
