package gomate

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// CodeRef is a reference to a specific piece of code in a textmate document
type CodeRef struct {
	AbsPath  string
	Line     int
	Column   int
	RelPath  string
	Filename string
	Excerpt  string
}

// URL returns a textmate specific URL pointing to the code ref
func (cf CodeRef) URL() *url.URL {
	return &url.URL{
		Scheme: "txmt",
		Path:   "open",
		RawQuery: strings.Join(
			[]string{
				fmt.Sprintf("line=%d", cf.Line),
				fmt.Sprintf("column=%d", cf.Column),
				fmt.Sprintf("url=file://%s", cf.AbsPath),
				fmt.Sprintf("title=%s", cf.RelPath),
			},
			"&",
		),
	}
}

// String returns the URL representation of a code ref
func (cf CodeRef) String() string {
	return cf.URL().String()
}

// CalcOffset will return the byte offset for the specified line and column of
// the reader (e.g. source code file)
func CalcOffset(r io.Reader, line, col uint) (uint, error) {
	var (
		s  = bufio.NewScanner(r)
		bc = uint(0)
	)

	for i := uint(1); s.Scan() && i < line; i++ {
		if err := s.Err(); err != nil {
			return 0, fmt.Errorf("unable to scan file for offset: %s", err)
		}

		bc += uint(len(s.Text()) + 1) // add newline to text length
	}

	return bc + col, nil
}
