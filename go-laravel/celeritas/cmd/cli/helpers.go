package main

import (
	"github.com/joho/godotenv"
	"os"
)

func setup() {
	err := godotenv.Load()
	if err != nil {
		exitGracefully(err)
	}

	// Note! This assumes you run the command from a folder that has a `.env` file
	path, err := os.Getwd()
	if err != nil {
		exitGracefully(err)
	}

	cel.RootPath = path
	cel.DB.DatabaseType = os.Getenv("DATABASE_TYPE")
}
