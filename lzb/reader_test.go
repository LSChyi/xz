package lzb

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestReader(t *testing.T) {
	filename := "fox.lzma"
	want := "The quick brown fox jumps over the lazy dog.\n"
	for i := 0; i < 2; i++ {
		f, err := os.Open(filename)
		if err != nil {
			t.Fatalf("os.Open(%q) error %s", filename, err)
		}
		p := make([]byte, 13)
		_, err = io.ReadFull(f, p)
		if err != nil {
			t.Fatalf("io.Readfull error %s", err)
		}
		params := Parameters{DictSize: 0x800000}
		params.SetProperties(Properties(p[0]))
		params.BufferSize = params.DictSize
		if i == 1 {
			params.SizeInHeader = true
			params.Size = int64(len(want))
		}
		r, err := NewReader(f, params)
		if err != nil {
			t.Fatalf("NewReader error %s", err)
		}
		buf := &bytes.Buffer{}
		if _, err = io.Copy(buf, r); err != nil {
			t.Fatalf("Copy error %s", err)
		}
		if err = f.Close(); err != nil {
			t.Fatalf("Close error %s", err)
		}
		got := buf.String()
		if got != want {
			t.Fatalf("read %q; but want %q", got, want)
		}
	}
}
