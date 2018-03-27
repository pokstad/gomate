package gomate

import (
	"fmt"
	"net/url"
	"strings"
)

type ExitCode int

const (
	ExitDiscard ExitCode = iota + 200
	ExitReplaceText
	ExitReplaceDocument
	ExitInsertText
	ExitInsertSnippet
	ExitShowHTML
	ExitShowToolTip
	ExitCreateNewDoc
	ExitInsertSnipNoIdent
)

// CodeRef is a reference to a specific piece of code in a textmate document
type CodeRef struct {
	AbsPath  string
	Line     uint
	Column   uint
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
