package html_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/pokstad/gomate/doc"
	"github.com/pokstad/gomate/html"
	"github.com/pokstad/gomate/test"
)

func TestGetDocHTML(t *testing.T) {
	b := new(bytes.Buffer)

	s := doc.Symbol{
		Name:   "Meow",
		Import: "github.com/pokstad/gomate/doc/testdata",
		Pkg:    "main",
		Decl:   "type Meow struct {\n\tLoudness uint   // how loud is this meow?\n\tSounds   string // how would you spell the sound of this meow?\n}",
		Doc:    "Meow is an important test identifer for kitties\n",
		Pos: doc.Position{
			AbsPath: test.MustGetAbsPath(t, "testdata/test.go"),
			Line:    0x4,
			Column:  0x6,
		},
	}

	err := html.SymbolDocHTML(b, os.Getenv("GOPATH"), s, nil)
	if err != nil {
		t.Fatalf("unable to render symbol: %s", err)
	}

	t.Logf("HTML output: %s", b.String())

}
