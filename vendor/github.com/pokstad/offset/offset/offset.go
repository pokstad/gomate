package offset

import (
	"errors"
	"io"
	"text/scanner"
)

var (
	ErrOffsetNotFound = errors.New("line and column offset not found")
)

// CalcRune will return the byte offset for the specified line and column of
// the reader (e.g. source code file). Line number will start counting at 1,
// while column starts counting at 0.
func CalcRune(r io.Reader, line, col int) (int, error) {
	var (
		s scanner.Scanner
	)

	s.Init(r)

	for {
		r := s.Next()

		if r == scanner.EOF {
			break
		}

		p := s.Pos()

		if p.Line == line && p.Column > col {
			break
		}

		if p.Line > line {
			break
		}

		if p.Line == line && p.Column == col {
			return p.Offset, nil
		}

	}

	return 0, ErrOffsetNotFound
}
