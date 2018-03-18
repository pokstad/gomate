package doc_test

import (
	"bytes"
	"go/doc"
	"go/parser"
	"go/token"
	"testing"
)

func TestDocHTML(t *testing.T) {
	pkgs, err := parser.ParseDir(
		token.NewFileSet(),
		"testdata",
		nil,
		parser.ParseComments,
	)
	if err != nil {
		t.Fatalf("unable to parse test file: %s", err)
	}

	for pName, pAST := range pkgs {
		pDoc := doc.New(pAST, "", doc.AllDecls)

		b := new(bytes.Buffer)
		doc.ToHTML(b, pDoc.Doc, nil)

		t.Logf("Package %s: %s", pName, b.String())

	}
}
