package user

import (
	"github.com/ATechnoHazard/gosh/utils"
	"log"
	"os"
	"strings"
)

type UserProfile struct {
	Username string
	Hostname string
	Path string
}

// SetupUserProfile initialise user profile variables
func (up *UserProfile) SetupUserProfile() {
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
	log.Println(dir)
	if path[0] == "home" && path[1] == username {
		path = utils.Remove(utils.Remove(path, 0), 0)
	}

	// replace homedir with ~
	path = utils.Unshift(path, "~")
	up.Path = strings.Join(path, "/")
}