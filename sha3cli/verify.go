package sha3cli

import (
	"bufio"
	"crypto/subtle"
	"fmt"
	"os"
	"strings"

	"ekyu.moe/util/cli"
)

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

	digest, err := computeFile(filename)
	if err != nil {
		if os.IsNotExist(err) && ignoreMissing {
			return 0
		}
		fmt.Fprintln(errWriter, err)
		fmt.Fprintln(outWriter, filename+": FAILED open or read")
		return 2
	}

	if subtle.ConstantTimeCompare(digest, checksum) == 1 {
		if !quiet {
			fmt.Fprintln(outWriter, filename+": OK")
		}
		return 0
	}

	fmt.Fprintln(outWriter, filename+": FAILED")
	return 1
}

func verifyChksumFile(chksumFilename string) ([4]int8, error) {
	errStat := [4]int8{0, 0, 0, 0}

	chksumFile, _, err := cli.AccessOpenFile(chksumFilename)
	if err != nil {
		return errStat, err
	}
	defer chksumFile.Close()

	scanner := bufio.NewScanner(chksumFile)

	for scanner.Scan() {
		line := scanner.Text()
		if errNo := verifyChecksumLine(line); errNo > 0 && errNo < 4 {
			errStat[errNo]++
		}
	}
	if err := scanner.Err(); err != nil {
		return errStat, err
	}

	return errStat, nil
}

func runVerify(filelist []string) int {
	var unmatchErrCount uint64 = 0
	var failReadErrCount uint64 = 0
	var wrongFormatErrCount uint64 = 0

	for _, filename := range filelist {
		errStat, err := verifyChksumFile(filename)
		if err != nil {
			fmt.Fprintln(errWriter, err)
			continue
		}
		unmatchErrCount += uint64(errStat[1])
		failReadErrCount += uint64(errStat[2])
		wrongFormatErrCount += uint64(errStat[3])
	}

	if unmatchErrCount+failReadErrCount+wrongFormatErrCount > 0 {
		if unmatchErrCount > 0 {
			fmt.Fprintf(errWriter, "WARNING: %d computed checksum(s) did NOT match\n", unmatchErrCount)
		}
		if failReadErrCount > 0 {
			fmt.Fprintf(errWriter, "WARNING: %d file(s) could not be read\n", failReadErrCount)
		}
		if wrongFormatErrCount > 0 {
			fmt.Fprintf(errWriter, "WARNING: %d line(s) is improperly formatted\n", wrongFormatErrCount)
		}

		return 1
	}

	return 0
}
