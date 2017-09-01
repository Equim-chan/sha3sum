package sha3cli

import (
	"fmt"
	"io"

	"ekyu.moe/util/cli"
)

func computeFile(filename string) ([]byte, error) {
	inFile, _, err := cli.AccessOpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	hasher := newHash()
	if _, err := io.Copy(hasher, inFile); err != nil {
		return nil, err
	}

	digest := hasher.Sum(nil)
	return digest, nil
}

func runCompute(filelist []string) int {
	for _, filename := range filelist {
		digest, err := computeFile(filename)
		if err != nil {
			fmt.Fprintln(errWriter, err)
			continue
		}
		fmt.Printf("%x  %s\n", digest, filename)
	}

	return 0
}
