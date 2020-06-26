package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case b >= 65 && b <= 90:
		if b < 78 {
			b += 13
		} else if b >= 78 {
			b -= 13
		}
	case b >= 97 && b <= 122:
		if b < 110 {
			b += 13
		} else if b >= 110 {
			b -= 13
		}
	}
	return b
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, _ := r.r.Read(b)

	for i, x := range b {
		b[i] = rot13(x)
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
