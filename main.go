package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/philippklemmer/getGithubUser/githubapi"
)

func parseArgs(args []string) (string, error) {
	if len(args) != 2 {
		if len(args) < 2 {
			return "", errors.New("You have to enter an Github Accountname to use this tool")
		} else {
			return "", errors.New("Enter only one Github Account as a Paramater")
		}
	}
	return args[1], nil
}

func main() {
	username, err := parseArgs(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Searching for a user with the name:", username)

	data, err := githubapi.GetGithubUser(username)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.Name, data.Login)
	fmt.Println(data.HtmlUrl)
}
