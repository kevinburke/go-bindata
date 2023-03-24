// This work is subject to the CC0 1.0 Universal (CC0 1.0) Public Domain Dedication
// license. Its contents can be found at:
// http://creativecommons.org/publicdomain/zero/1.0/

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	AppVersion = "1.1.0"
)

func main() {
	msg := `This version of go-bindata is not functional and likely running because you ran

	go get github.com/kevinburke/go-bindata@latest

with no major version selector. 

Try amending the command to:

	go get github.com/kevinburke/go-bindata/v4/...@latest

Note the "v4" and the "..." - you will need both to retrieve the latest version.
Apologies for the confusion, this is unavoidable behavior from the Go tool.
`
	flag.Usage = func() { fmt.Println(msg) }
	flag.Parse()
	if flag.Arg(1) == "version" {
		fmt.Printf("go-bindata %s (deprecated)\n", AppVersion)
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, msg)
	os.Exit(1)
}
