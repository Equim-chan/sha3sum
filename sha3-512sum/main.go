package main

import (
	"os"

	"ekyu.moe/sha3sum/sha3cli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := sha3cli.Run(sha3.New512, 512)
	os.Exit(exitCode)
}
