package gomate

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"net/url"
	"strings"
)

const errWarnInfoTmpl = `
<html>
<head>
</head>
<body>
{{ range $category, $coderefs := . }}
<div>
<h2> {{ $category }} </h2>
<ul>
{{ range $coderefs }}
<li><a href="{{ print . | safeURL }}">{{ .Filename }}:{{ .Line }}:{{ .Column }}:</a> {{ .Excerpt }}</li>
{{ end }}
</ul>
</div>
{{ end }}
</body></html>
`

var htmlTmpl = template.Must(template.New("").Funcs(
	template.FuncMap{
		"safeURL": func(u string) template.URL { return template.URL(u) },
	},
).Parse(errWarnInfoTmpl))

// RenderHTML will render error, warning, and info code references
func RenderHTML(w io.Writer, errors, warnings, infos []CodeRef) error {
	return htmlTmpl.Execute(w, map[string][]CodeRef{
		"errors":   errors,
		"warnings": warnings,
		"infos":    infos,
	})
}

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
