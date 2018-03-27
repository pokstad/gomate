package html

import (
	"html/template"
	"io"

	"github.com/pokstad/gomate"
)

const codeRefsTmpl = `
<html>
	<head>
		<title>{{ .Title }}</title>
		<style type="text/css">
			{{ .Stylesheet }}
		</style>
	</head>
	<body class="remarkdown">
		<div>
			<h1> {{ .Title }} </h1>
			<ul>
				{{ range .Refs }}
					<li><a href="{{ print . | safeURL }}">{{ .Filename }}:{{ .Line }}:{{ .Column }}:</a> {{ .Excerpt }}</li>
				{{ end }}
			</ul>
		</div>
	</body>
</html>
`

var htmlTmpl = template.Must(template.New("").Funcs(tmplFuncs).Parse(codeRefsTmpl))

// CodeRefsHTML will render a list of code references
func CodeRefsHTML(w io.Writer, title string, refs []gomate.CodeRef, css []byte) error {
	return htmlTmpl.Execute(w, struct {
		Title      string
		Refs       []gomate.CodeRef
		Stylesheet template.CSS
	}{
		Title:      title,
		Refs:       refs,
		Stylesheet: template.CSS(string(css)),
	})
}
