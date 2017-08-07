package main

import (
	"fmt"
)

func timingSafeEqual(a, b []byte) bool {
	size := len(a)
	if size != len(b) {
		return false
	}

	var x byte = 0
	for i := 0; i < size; i++ {
		x |= a[i] ^ b[i]
	}

	return x == 0
}

func printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(outWriter, format, a...)
}

func println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(outWriter, a...)
}

func printError(e error) (n int, err error) {
	return fmt.Fprintln(errWriter, TITLE+": "+e.Error())
}
