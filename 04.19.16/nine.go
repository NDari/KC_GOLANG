package main

import (
	"bytes"
	"io"
)

func main() {
	var buf *bytes.Buffer
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("Done!\n"))
	}
}
