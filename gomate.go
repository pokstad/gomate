package gomate

import (
	"fmt"
	"net/url"
	"strings"
)

// ExitCode is an enumeration of the different exit codes supported by Textmate.
// Note: exit codes are ignored if HTML output is being used
// Refer to the official docs:
// https://github.com/textmate/bundle-support.tmbundle/blob/master/Support/shared/lib/exit_codes.rb
type ExitCode int

const (
	// ExitDiscard will discard output
	ExitDiscard ExitCode = iota + 200
	// ExitReplaceText will replace the selected text with the output
	ExitReplaceText
	// ExitReplaceDocument will replace the entire document with the output
	ExitReplaceDocument
	// ExitInsertText will insert the output at the cursor location
	ExitInsertText
	// ExitInsertSnippet will insert the snippet at the cursor location
	ExitInsertSnippet
	// ExitShowHTML will render output as HTML in a separate window
	ExitShowHTML
	// ExitShowToolTip will show the output as a tool tip near the cursor
	ExitShowToolTip
	// ExitCreateNewDoc will create a new document with the output
	ExitCreateNewDoc
	// ExitInsertSnipNoIdent will insert the snippet at the cursor with no
	// additional indentation
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
