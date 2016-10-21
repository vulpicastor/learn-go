/*
Implements methods/23 of A Tour of Go
*/
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(b []byte) (n int, err error) {
	inter := make([]byte, len(b))
	n, err = r.r.Read(inter)
	for i, v := range inter[:n] {
		if ('A' <= v) && (v <= 'Z') {
			inter[i] = (v-'4')%26 + 'A'
		} else if ('a' <= v) && (v <= 'z') {
			inter[i] = (v-'T')%26 + 'a'
		}
	}
	n = copy(b, inter[:n])
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
