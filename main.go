package main

import (
	"github.com/joho/godotenv"
	"github.com/rbrick/bytes/cmd"
)

func init() {
	// Load the .env file, if any.
	godotenv.Load()
}

func main() {
	cmd.Execute()
}
