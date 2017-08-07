# sha3sum
[![Travis](https://img.shields.io/travis/Equim-chan/sha3sum.svg)](https://travis-ci.org/Equim-chan/sha3sum)
[![Go Report Card](https://goreportcard.com/badge/github.com/Equim-chan/sha3sum)](https://goreportcard.com/report/github.com/Equim-chan/sha3sum)
[![license](https://img.shields.io/badge/BSD-3.0-blue.svg)](https://github.com/Equim-chan/sha3sum/blob/master/LICENSE)

A sha3sum CLI utility based on [golang.org/x/crypto/sha3](https://godoc.org/golang.org/x/crypto/sha3)

## Install
Quick install:
```bash
$ go get -u ekyu.moe/sha3sum
$ go generate ekyu.moe/sha3sum
$ $GOPATH/bin/sha3-224sum --help
```

Build binaries for all platforms using gox:
```bash
$ go get -u github.com/mitchellh/gox
$ BUILD_ALL=1 go generate ekyu.moe/sha3sum
```

## License
[BSD-3.0](https://github.com/Equim-chan/sha3sum/blob/master/LICENSE)
