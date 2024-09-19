package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const dockerHost = "tcp://localhost:2375"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: envdocker [--] <command> [args...]")
		fmt.Println("Executes the given command with DOCKER_HOST environment variable set.")
		os.Exit(1)
	}

	// Remove the "--" separator if present
	args := os.Args[1:]
	if args[0] == "--" {
		args = args[1:]
	}

	if len(args) == 0 {
		fmt.Println("Error: No command specified")
		os.Exit(1)
	}

	// Check if the command exists
	cmdPath, err := exec.LookPath(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "envdocker: Error: Command '%s' not found in PATH\n", args[0])
		fmt.Fprintf(os.Stderr, "Please make sure the command is installed and accessible.\n")
		os.Exit(1)
	}

	cmd := exec.Command(cmdPath, args[1:]...)

	env := os.Environ()
	env = append(env, fmt.Sprintf("DOCKER_HOST=%s", dockerHost))
	cmd.Env = env

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("envdocker: Executing with DOCKER_HOST=%s\n", dockerHost)
	fmt.Printf("Command: %s\n", strings.Join(args, " "))

	err = cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "envdocker: Error executing command: %v\n", err)
		os.Exit(1)
	}
}
