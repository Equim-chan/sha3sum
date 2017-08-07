package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func runCompute(args []string) int {
	for _, arg := range args {
		if arg == "-" {
			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				printError(err)
				continue
			}
			digest := CHECKSUM(data, outputSize)
			printf("%x  -\n", digest)
			continue
		}

		matches, err := filepath.Glob(arg)
		if err != nil {
			printError(err)
			continue
		}
		for _, filename := range matches {
			data, err := ioutil.ReadFile(filename)
			if err != nil {
				printError(err)
				continue
			}
			digest := CHECKSUM(data, outputSize)
			printf("%x  %s\n", digest, filename)
		}
	}

	return 0
}
