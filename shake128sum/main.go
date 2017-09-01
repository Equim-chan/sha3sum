package main

import (
	"os"

	"ekyu.moe/sha3sum/shakecli"
	"golang.org/x/crypto/sha3"
)

func main() {
	exitCode := shakecli.Run(sha3.NewShake128, 128)
	os.Exit(exitCode)
}
