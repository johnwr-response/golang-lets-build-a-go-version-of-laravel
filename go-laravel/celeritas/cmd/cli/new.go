package main

import (
	"log"
	"strings"
)

func doNew(args2 string) {
	appName := strings.ToLower(args2)
	log.Println("appName: ", appName)

	// sanitize the application name (convert url to single word)

	// git clone the skeleton application

	// remove the .git directory

	// create a ready to go .env file

	// create a makefile

	// update the go.mod file

	// update the existing .go files with correct name/imports

	// run go mod tidy in the project directory
}
