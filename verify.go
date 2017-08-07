package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func runVerify(args []string) int {
	var unmatchErrCount uint64 = 0
	var failReadErrCount uint64 = 0
	var wrongFormatErrCount uint64 = 0

	for _, arg := range args {
		if arg == "-" {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				line := scanner.Text()
				switch verifyChecksumLine(line) {
				case 1:
					unmatchErrCount++
				case 2:
					failReadErrCount++
				case 3:
					wrongFormatErrCount++
				}
			}
			if err := scanner.Err(); err != nil {
				printError(err)
			}
			continue
		}

		matches, err := filepath.Glob(arg)
		if err != nil {
			printError(err)
			continue
		}
		for _, filename := range matches {
			content, err := ioutil.ReadFile(filename)
			if err != nil {
				printError(err)
				continue
			}
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				switch verifyChecksumLine(line) {
				case 1:
					unmatchErrCount++
				case 2:
					failReadErrCount++
				case 3:
					wrongFormatErrCount++
				}
			}
		}
	}

	if unmatchErrCount+failReadErrCount+wrongFormatErrCount > 0 {
		if unmatchErrCount == 1 {
			printError(fmt.Errorf("WARNING: %d computed checksum did NOT match", unmatchErrCount))
		} else if unmatchErrCount > 1 {
			printError(fmt.Errorf("WARNING: %d computed checksums did NOT match", unmatchErrCount))
		}
		if failReadErrCount == 1 {
			printError(fmt.Errorf("WARNING: %d file could not be read", failReadErrCount))
		} else if failReadErrCount > 1 {
			printError(fmt.Errorf("WARNING: %d files could not be read", failReadErrCount))
		}
		if wrongFormatErrCount == 1 {
			printError(fmt.Errorf("WARNING: %d line is improperly formatted", wrongFormatErrCount))
		} else if wrongFormatErrCount > 1 {
			printError(fmt.Errorf("WARNING: %d lines are improperly formatted", wrongFormatErrCount))
		}
		return 1
	}

	return 0
}

func verifyChecksumLine(line string) int8 {
	if strings.Trim(line, " ") == "" {
		return -1
	}

	var checksum []byte
	var filename string

	_, err := fmt.Sscanf(line, "%x%s", &checksum, &filename)
	if err != nil {
		return 3
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		if ignoreMissing {
			return 0
		}
		printError(err)
		println(filename + ": FAILED open or read")
		return 2
	}

	digest := CHECKSUM(data, outputSize)
	if timingSafeEqual(digest[:], checksum) {
		if !quiet {
			println(filename + ": OK")
		}
		return 0
	}

	println(filename + ": FAILED")
	return 1
}
