package shakecli

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"ekyu.moe/util/cli"
	"golang.org/x/crypto/sha3"
)

var (
	newHash func() sha3.ShakeHash

	title    string
	algoName string

	checkMode     bool
	ignoreMissing bool
	quiet         bool
	status        bool
	outputSize    uint

	outWriter io.Writer
	errWriter io.Writer
)

func Run(hfunc func() sha3.ShakeHash, bits int) int {
	newHash = hfunc
	title = "shake" + strconv.Itoa(bits) + "sum"
	algoName = "SHAKE-" + strconv.Itoa(bits)

	flag.BoolVar(&checkMode, "c", false, "")
	flag.BoolVar(&checkMode, "check", false, "")
	flag.BoolVar(&ignoreMissing, "ignore-missing", false, "")
	flag.BoolVar(&quiet, "quiet", false, "")
	flag.BoolVar(&status, "status", false, "")
	flag.UintVar(&outputSize, "size", 32, "")
	flag.UintVar(&outputSize, "s", 32, "")

	usage := "Usage: " + title + ` [OPTION]... [FILE]...
Print or check ` + algoName + ` checksums.
With no FILE, or when FILE is -, read standard input.
  -s, --size           length of output (bytes) for ` + algoName + `, default 32
  -c, --check          read ` + algoName + ` sums from the FILEs and check them

The following three options are useful only when verifying checksums:
      --ignore-missing  don't fail or report status for missing files
      --quiet          don't print OK for each successfully verified file
      --status         don't output anything, status code shows success

  -h, --help     display this help and exit

The sums are computed as described in FIPS-202.
When checking, the input should be a former output of this program.
The default mode is to print a line with checksum, a space,
and name for each FILE.
`

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	flag.Parse()

	if status && checkMode {
		outWriter = ioutil.Discard
		errWriter = ioutil.Discard
	} else {
		outWriter = os.Stdout
		errWriter = os.Stderr
	}

	filelist, _ := cli.ParseFileList(flag.Args(), true)

	if checkMode {
		return runVerify(filelist)
	} else {
		return runCompute(filelist, outputSize)
	}
}
