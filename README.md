# sha3sum
[![Travis](https://img.shields.io/travis/Equim-chan/sha3sum.svg)](https://travis-ci.org/Equim-chan/sha3sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/sha3sum)](https://goreportcard.com/report/github.com/Equim-chan/sha3sum)
[![license](https://img.shields.io/badge/BSD-3.0-blue.svg)](https://github.com/Equim-chan/sha3sum/blob/master/LICENSE)

A sha3sum CLI utility based on [golang.org/x/crypto/sha3](https://godoc.org/golang.org/x/crypto/sha3)

It covers the following utilities:

* sha3-224sum
* sha3-256sum
* sha3-384sum
* sha3-512sum
* shake128sum
* shake256sum

## Install
Quick install:
```bash
$ go get -d -u ekyu.moe/sha3sum/...
$ go generate ekyu.moe/sha3sum
$ echo -n "Hello, 世界" | sha3-224sum
ee346b66418f901d68c35fc02d25d5a3cf8ee0fcea32c3ded16b82d0  -
```

Build binaries for all platforms using gox:
```bash
$ go get -u github.com/mitchellh/gox
$ BUILD_ALL=1 go generate ekyu.moe/sha3sum
```

Also, you may check the [release page](https://github.com/Equim-chan/sha3sum/releases) for handy binaries.

## Help
```bash
$ sha128sum --help
Usage: shake128sum [OPTION]... [FILE]...
Print or check SHAKE-128 checksums.
With no FILE, or when FILE is -, read standard input.

  -s, --size           length of output (bytes) for SHAKE-128, default 32
  -c, --check          read SHAKE-128 sums from the FILEs and check them

The following three options are useful only when verifying checksums:
      --ignore-missing  don't fail or report status for missing files
      --quiet          don't print OK for each successfully verified file
      --status         don't output anything, status code shows success

  -h, --help     display this help and exit
      --version  output version information and exit

The sums are computed as described in FIPS-180-4.
When checking, the input should be a former output of this program.
The default mode is to print a line with checksum, a space,
and name for each FILE.
```

`-s` flag is only available for `shake128sum` and `shake256sum`.

## License
[BSD-3.0](https://github.com/Equim-chan/sha3sum/blob/master/LICENSE)
