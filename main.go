package main

import (
		"fmt"
		"os"

//		"github.com/philippklemmer/getGithubUser/githubapi"
)

func main() {
fmt.Print(os.Args)
	if len(os.Args != 2 {
		if len(os.Args < 2) {
			fmt.Println("You have to enter an Github Accountname to use this tool")
			} else {
		fmt.Printkn("Enter only one Github Account as a Parameter")
		os.Exit()
	}

	username := os.Args[1]
//user := githubapi.getGithubUser(name)	
}
