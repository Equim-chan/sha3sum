package main

//go:generate go run $GOPATH/src/ekyu.moe/sha3sum/build_tool/main.go

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/sha3"
)

var (
	TITLE     string
	ALGO_NAME string
	VERSION   string
	CHECKSUM  func([]byte, uint) []byte

	flagSet       *flag.FlagSet
	checkMode     bool
	ignoreMissing bool
	quiet         bool
	status        bool
	printVersion  bool
	outputSize    uint

	outWriter io.Writer
	errWriter io.Writer
)

func init() {
	switch ALGO_NAME {
	case "SHA3-224":
		CHECKSUM = func(data []byte, _ uint) []byte {
			digest := sha3.Sum224(data)
			return digest[:]
		}
	case "SHA3-256":
		CHECKSUM = func(data []byte, _ uint) []byte {
			digest := sha3.Sum256(data)
			return digest[:]
		}
	case "SHA3-384":
		CHECKSUM = func(data []byte, _ uint) []byte {
			digest := sha3.Sum384(data)
			return digest[:]
		}
	case "SHA3-512":
		CHECKSUM = func(data []byte, _ uint) []byte {
			digest := sha3.Sum512(data)
			return digest[:]
		}
	case "SHAKE-128":
		CHECKSUM = func(data []byte, outLen uint) []byte {
			digest := make([]byte, outLen)
			sha3.ShakeSum128(digest, data)
			return digest
		}
	case "SHAKE-256":
		CHECKSUM = func(data []byte, outLen uint) []byte {
			digest := make([]byte, outLen)
			sha3.ShakeSum256(digest, data)
			return digest
		}
	}

	flagSet = flag.NewFlagSet(TITLE, flag.ContinueOnError)

	flagSet.BoolVar(&checkMode, "check", false, "")
	flagSet.BoolVar(&checkMode, "c", false, "")
	flagSet.BoolVar(&ignoreMissing, "ignore-missing", false, "")
	flagSet.BoolVar(&quiet, "quiet", false, "")
	flagSet.BoolVar(&status, "status", false, "")
	flagSet.BoolVar(&printVersion, "version", false, "")

	usage := "Usage: " + TITLE + ` [OPTION]... [FILE]...
Print or check ` + ALGO_NAME + ` checksums.
With no FILE, or when FILE is -, read standard input.` + "\n\n"

	if ALGO_NAME == "SHAKE-128" || ALGO_NAME == "SHAKE-256" {
		flagSet.UintVar(&outputSize, "size", 32, "")
		flagSet.UintVar(&outputSize, "s", 32, "")
		usage += "  -s, --size           length of output (bytes) for " + ALGO_NAME + ", default 32\n"
	}

	usage += `  -c, --check          read ` + ALGO_NAME + ` sums from the FILEs and check them

The following three options are useful only when verifying checksums:
      --ignore-missing  don't fail or report status for missing files
      --quiet          don't print OK for each successfully verified file
      --status         don't output anything, status code shows success

  -h, --help     display this help and exit
      --version  output version information and exit

The sums are computed as described in FIPS-180-4.
When checking, the input should be a former output of this program.
The default mode is to print a line with checksum, a space,
and name for each FILE.`

	flagSet.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}
}

func main() {
	flagSet.Parse(os.Args[1:])

	if printVersion {
		fmt.Fprintln(os.Stderr, TITLE+" "+VERSION)
		os.Exit(2)
	}

	if status && checkMode {
		outWriter = ioutil.Discard
		errWriter = ioutil.Discard
	} else {
		outWriter = os.Stdout
		errWriter = os.Stderr
	}

	args := flagSet.Args()
	if len(args) == 0 {
		args = append(args, "-")
	}

	exitCode := 0
	if checkMode {
		exitCode = runVerify(args)
	} else {
		exitCode = runCompute(args)
	}

	os.Exit(exitCode)
}
