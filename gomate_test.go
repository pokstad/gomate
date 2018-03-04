package gomate_test

import (
	"bytes"
	"testing"

	"github.com/pokstad/gomate"
)

func TestCalcOffset(t *testing.T) {
	contents := []byte("line 1\nline 2\nline 3")
	b := bytes.NewBuffer(contents)

	o, err := gomate.CalcOffset(b, 3, 2)
	if err != nil {
		t.Fatalf("cannot calculate offset: %s", err)
	}

	if o != 16 {
		t.Fatalf("unexpected rune count: %d", o)
	}
}
