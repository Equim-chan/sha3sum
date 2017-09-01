package sha3cli // import "ekyu.moe/sha3sum/sha3cli"

import (
	"flag"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"ekyu.moe/util/cli"
)

var (
	newHash func() hash.Hash

	title    string
	algoName string

	checkMode     bool
	ignoreMissing bool
	quiet         bool
	status        bool

	outWriter io.Writer
	errWriter io.Writer
)

func Run(hfunc func() hash.Hash, bits int) int {
	newHash = hfunc
	title = "sha3-" + strconv.Itoa(bits) + "sum"
	algoName = "SHA3-" + strconv.Itoa(bits)

	flag.BoolVar(&checkMode, "c", false, "")
	flag.BoolVar(&checkMode, "check", false, "")
	flag.BoolVar(&ignoreMissing, "ignore-missing", false, "")
	flag.BoolVar(&quiet, "quiet", false, "")
	flag.BoolVar(&status, "status", false, "")

	usage := "Usage: " + title + ` [OPTION]... [FILE]...
Print or check ` + algoName + ` checksums.
With no FILE, or when FILE is -, read standard input.

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
		return runCompute(filelist)
	}
}
