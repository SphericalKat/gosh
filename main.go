package main

import (
	"fmt"
	"github.com/ATechnoHazard/gosh/user"
	"github.com/ATechnoHazard/gosh/utils"
	"github.com/eiannone/keyboard"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Struct to hold user profile vars loaded from the environment

func main() {
	// create a reader object to read data from stdin
	//reader := bufio.NewReader(os.Stdin)
	err := keyboard.Open()
	if err != nil {
		log.Panic(err)
	}
	defer keyboard.Close()
	var input string

	// create a channel to notify us about SIGINT and SIGTERMs
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	profile := new(user.UserProfile)
	profile.SetupUserProfile()
	fmt.Printf("%s@%s:%s$ ", profile.Username, profile.Hostname, profile.Path)

	// main input loop
	for {
		profile.SetupUserProfile()

		// OS signal was sent
		select {
		case sig := <-sigs:
			fmt.Println(sig)
			if sig == syscall.SIGQUIT || sig == syscall.SIGINT {
				continue
			}

		// regular input
		default:
			char, key, err := keyboard.GetKey()
			if err != nil {
				log.Panic(err)
			}

			switch key {
			case keyboard.KeyEnter:
				fmt.Print("\n")
				// execute input command and log errors if any
				if err = utils.ExecInput(input); err != nil {
					_, err := fmt.Fprintln(os.Stderr, err)
					if err != nil {
						panic(err)
					}
				}
				fmt.Printf("%s@%s:%s$ ", profile.Username, profile.Hostname, profile.Path)
				input = ""
				break
			case keyboard.KeySpace:
				input += " "
				fmt.Print(" ")
			default:
				input += string(char)
				fmt.Print(string(char))
			}
		}
	}
}
