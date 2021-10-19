package main

import (
	"fmt"
	"os"
	"strings"

	colour "github.com/fatih/color"
	flag "github.com/ogier/pflag"
)

// flags
var (
	user       string
	repository string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		noFlags()
	}

	var users = strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for _, u := range users {
		result := GetUsers(u)
		colour.Cyan(`Username:	%s`, result.Login)
		colour.Blue(`Name:		%s`, result.Name)
		colour.Green(`Email:		%s`, result.Email)
		colour.HiMagenta(`Bio:		%s`, result.Bio)
		fmt.Println("")
	}
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
	flag.StringVarP(&repository, "repository", "r", "", "GitHub repository")
}

func noFlags() {
	fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
