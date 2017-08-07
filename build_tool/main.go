package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const (
	VERSION = "0.1.0"
)

var (
	target = [...][2]string{
		[2]string{"sha3-224sum", "SHA3-224"},
		[2]string{"sha3-256sum", "SHA3-256"},
		[2]string{"sha3-384sum", "SHA3-384"},
		[2]string{"sha3-512sum", "SHA3-512"},
		[2]string{"shake128sum", "SHAKE-128"},
		[2]string{"shake256sum", "SHAKE-256"},
	}
)

func main() {
	GOPATH, ok := os.LookupEnv("GOPATH")
	if !ok {
		fmt.Fprintln(os.Stderr, "build error: $GOPATH not found")
		os.Exit(1)
	}

	_, buildAll := os.LookupEnv("BUILD_ALL")
	for _, item := range target {
		var cmd *exec.Cmd
		if buildAll {
			args := []string{
				"-output", GOPATH + "/bin/{{.OS}}_{{.Arch}}/" + item[0],
				"-ldflags",
				"-X main.TITLE=" + item[0] + " -X main.ALGO_NAME=" + item[1] + " -X main.VERSION=" + VERSION + " -s -w",
				"ekyu.moe/sha3sum",
			}
			cmd = exec.Command("gox", args...)
		} else {
			GOEXE := ""
			if runtime.GOOS == "windows" {
				GOEXE = ".exe"
			}
			args := []string{
				"build", "-o", GOPATH + "/bin/" + item[0] + GOEXE,
				"-ldflags",
				"-X main.TITLE=" + item[0] + " -X main.ALGO_NAME=" + item[1] + " -X main.VERSION=" + VERSION + " -s -w",
				"ekyu.moe/sha3sum",
			}
			cmd = exec.Command("go", args...)
		}
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(os.Stderr, "build error: "+err.Error())
		}
		if len(stdoutStderr) > 0 {
			fmt.Println(string(stdoutStderr))
		}
	}
}
