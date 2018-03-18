package html

import (
	"html/template"
	"io"
	"strings"

	"github.com/pokstad/gomate/doc"
)

const getdocTmpl = `
<html>
	<head>
		<title>{{ .Title }}</title>
		<style type="text/css">
			{{ .Stylesheet }}
		</style>
	</head>
	<body class="remarkdown">
	{{ $baseDir := .BaseDir }}
	{{ with .Symbol }}
		<div>
			<h1>
				<a href="{{ symbolRef . $baseDir | safeURL }}">
					"{{ .Import }}".{{ .Pkg }}.{{ .Name }}
				</a>
			</h1>
			<p>
				<pre>{{ .Decl }}</pre>	
			</p>
			<div>{{ toHTML . }}</div>
		</div>
	{{ end }}
	<hr/>
	<small>
		<p>Stderr:</p>
		<blockquote>
	</body>
</html>
`

var tmplFuncs = template.FuncMap{
	"safeURL":    func(u string) template.URL { return template.URL(u) },
	"breaklines": func(s string) template.HTML { return template.HTML(strings.Replace(s, "\n", "<br/>", -1)) },
	"toHTML":     func(s doc.Symbol) template.HTML { return template.HTML(s.HTML()) },
	"symbolRef":  func(s doc.Symbol, baseDir string) string { return s.CodeRef(baseDir).URL().String() },
}

var getdocHTMLTmpl = template.Must(template.New("").Funcs(tmplFuncs).Parse(getdocTmpl))

// SymbolDocHTML will render a symbol's documentation
func SymbolDocHTML(w io.Writer, baseDir string, s doc.Symbol, css []byte) error {
	return getdocHTMLTmpl.Execute(w, struct {
		Title      string
		BaseDir    string
		Stylesheet template.CSS
		Symbol     doc.Symbol
	}{
		Title:      s.Name,
		BaseDir:    baseDir,
		Stylesheet: template.CSS(string(css)),
		Symbol:     s,
	})
}
