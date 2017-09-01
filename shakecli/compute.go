package shakecli

import (
	"fmt"
	"io"

	"ekyu.moe/util/cli"
)

func computeFile(filename string, size uint) ([]byte, error) {
	inFile, _, err := cli.AccessOpenFile(filename)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()

	hasher := newHash()
	if _, err := io.Copy(hasher, inFile); err != nil {
		return nil, err
	}

	digest := make([]byte, size)
	if _, err := hasher.Read(digest); err != nil {
		return nil, err
	}

	return digest, nil
}

func runCompute(filelist []string, size uint) int {
	for _, filename := range filelist {
		digest, err := computeFile(filename, size)
		if err != nil {
			fmt.Fprintln(errWriter, err)
			continue
		}
		fmt.Printf("%x  %s\n", digest, filename)
	}

	return 0
}
