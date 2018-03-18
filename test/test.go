package test

import (
	"path/filepath"
	"testing"
)

func MustGetAbsPath(t *testing.T, relP string) string {
	p, err := filepath.Abs(relP)
	if err != nil {
		t.Fatalf("unable to obtain absolute path: %s", err)
	}
	return p
}
