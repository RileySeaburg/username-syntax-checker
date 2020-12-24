package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func main() {
	var userNameInput string
	flag.StringVar(&userNameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker")
	fmt.Println("Checking syntac for username, \"", userNameInput, "\", resulted in: ", CheckUsernameSyntax(userNameInput), "\n")
}

func CheckUsernameSyntax(username string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(username)
	return validationResult
}
