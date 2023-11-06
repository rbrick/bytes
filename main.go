package main

import (
	"github.com/joho/godotenv"
	"github.com/rbrick/bytes/cmd"
)

const (
	Bytes2Address = "0xa19f5264F7D7Be11c451C093D8f92592820Bea86"
	USDCAddress   = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
)

func init() {
	// Load the .env file, if any.
	godotenv.Load()
}

func main() {
	cmd.Execute()
}
