package main

import (
	"os"

	"ekyu.moe/sha3sum.v2/sha3cli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := sha3cli.Run(sha3.New224, 224)
	os.Exit(exitCode)
}
