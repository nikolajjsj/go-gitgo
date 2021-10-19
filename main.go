package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// flags
var (
	user       string
	repository string
)

func main() {
	flag.Parse()

	hasFlags()

	var users = strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
	flag.StringVarP(&repository, "repository", "r", "", "GitHub repository")
}

func hasFlags() {
	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
