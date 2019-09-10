package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ATechnoHazard/gosh/utils"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

// Struct to hold user profile vars loaded from the environment
type UserProfile struct {
	Username string
	Hostname string
	Path string
}

func main() {
	// create a reader object to read data from stdin
	reader := bufio.NewReader(os.Stdin)

	// create a channel to notify us about SIGINT and SIGTERMs
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	profile := new(UserProfile)


	// main input loop
	for {
		profile.setupUserProfile()
		fmt.Printf("%s@%s:%s$ ", profile.Username, profile.Hostname, profile.Path)

		// OS signal was sent
		select {
			case sig := <- sigs:
				fmt.Println(sig)
				if sig == syscall.SIGQUIT || sig == syscall.SIGINT {
					continue
				}

		// regular input
		default:
			input, err := reader.ReadString('\n')
			if err != nil {
				_, err = fmt.Fprintln(os.Stderr, err)
				if err != nil {
					panic(err)
				}
			}

			// execute input command and log errors if any
			if err = execInput(input); err != nil {
				_, err := fmt.Fprintln(os.Stderr, err)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

// initialise user profile variables
func (up *UserProfile) setupUserProfile() {
	username, ok := os.LookupEnv("USERNAME")
	if !ok {
		up.Username = "user"
	}
	up.Username = username

	hostname, ok := os.LookupEnv("HOSTNAME")
	if !ok {
		up.Hostname = "hostname"
	}
	up.Hostname = hostname

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	dir = "/home/amole/go/src/github.com/ATechnoHazard/gosh"

	// remove trailing homedir if present
	path := utils.Remove(strings.Split(dir, "/"), 0)
	if path[0] == "home" && path[1] == username {
		path = utils.Remove(utils.Remove(path, 0), 0)
	}

	// replace homedir with ~
	path = utils.Unshift(path, "~")
	up.Path = strings.Join(path, "/")
}

// utility function for executing input commands
func execInput(input string) error {
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
			return  errors.New("path required")
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

