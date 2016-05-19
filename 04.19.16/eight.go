package main

import "io"

func main() {
	var buf io.Writer
	f(buf)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("Done!\n"))
	}
}
