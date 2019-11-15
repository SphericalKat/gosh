package main

import (
	"bufio"
	"fmt"
	"github.com/ATechnoHazard/gosh/user"
	"github.com/ATechnoHazard/gosh/utils"
	"github.com/eiannone/keyboard"
	"os"
	"os/signal"
	"syscall"
)

// Struct to hold user profile vars loaded from the environment

func main() {
	// create a reader object to read data from stdin
	reader := bufio.NewReader(os.Stdin)

	// create a channel to notify us about SIGINT and SIGTERMs
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	profile := new(user.UserProfile)

	// main input loop
	for {
		profile.SetupUserProfile()
		fmt.Printf("%s@%s:%s$ ", profile.Username, profile.Hostname, profile.Path)

		// OS signal was sent
		select {
		case sig := <-sigs:
			fmt.Println(sig)
			if sig == syscall.SIGQUIT || sig == syscall.SIGINT {
				continue
			}

		// regular input
		default:
			//input, err := reader.ReadString('\n')
			//if err != nil {
			//	_, err = fmt.Fprintln(os.Stderr, err)
			//	if err != nil {
			//		panic(err)
			//	}
			//}



			// execute input command and log errors if any
			if err = utils.ExecInput(input); err != nil {
				_, err := fmt.Fprintln(os.Stderr, err)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
