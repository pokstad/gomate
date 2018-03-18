package test

import (
	"path/filepath"
	"testing"
)

// MustGetAbsPath will fail a test if not able to obtain a relative path of relP
func MustGetAbsPath(t *testing.T, relP string) string {
	p, err := filepath.Abs(relP)
	if err != nil {
		t.Fatalf("unable to obtain absolute path: %s", err)
	}
	return p
}
