package utils

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func Remove(a []string, i int) []string {
	return append(a[:i], a[i+1:]...)
}

func Unshift(a []string, b string) []string {
	return append([]string{b}, a...)
}

// ExecInput utility function for executing input commands
func ExecInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// special cases
	switch args[0] {
	case "":
		return nil
	case "exit":
		os.Exit(0)
		return nil
	case "quit":
		os.Exit(0)
		return nil
	case "logout":
		os.Exit(0)
		return nil
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	default:
		// Prepare the command to execute.
		cmd := exec.Command(args[0], args[1:]...)

		// Set the correct output device.
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout

		// Execute the command and return the error if any.
		return cmd.Run()
	}
}
