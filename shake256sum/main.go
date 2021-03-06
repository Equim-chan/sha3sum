package main

import (
	"os"

	"ekyu.moe/sha3sum/shakecli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := shakecli.Run(sha3.NewShake256, 256)
	os.Exit(exitCode)
}
