package main

import (
	"os"

	"ekyu.moe/sha3sum/sha3cli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := sha3cli.Run(sha3.New256, 256)
	os.Exit(exitCode)
}
