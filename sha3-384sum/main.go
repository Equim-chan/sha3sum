package main

import (
	"os"

	"ekyu.moe/sha3sum/sha3cli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := sha3cli.Run(sha3.New384, 384)
	os.Exit(exitCode)
}
