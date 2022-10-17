package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

type Rot13Reader struct {
	r io.Reader
	s string
}

func (reader Rot13Reader) Read(bytes []byte) (int, error) {
	var n, err = reader.r.Read(bytes)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {

		var b byte = bytes[i]

		if b >= 'A' && b <= 'Z' {
			bytes[i] = 'A' + (b-'A'+13)%26
		} else if b >= 'a' && b <= 'z' {
			bytes[i] = 'a' + (b-'a'+13)%26
		}
	}

	return n, nil
}

func TestSomeReader(t *testing.T) {
	sreader := strings.NewReader("Lbh penpxrq gur pbqr!")
	s := "some string"
	var r = Rot13Reader{s: s, r: sreader}
	io.Copy(os.Stdout, &r)
}
