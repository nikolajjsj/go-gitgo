package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

// flags
var (
	user string
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var users = strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
}

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}
