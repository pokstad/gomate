package html

import (
	"fmt"
	"html/template"
	"io"

	"github.com/pkg/errors"
)

const errTmpl = `
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
			<p>Error: {{ .ErrorSimple }}</p>
			<pre>{{ .ErrorDetail }}</pre>
		</div>
		<div>
			Logs:
			<small><blockquote><pre>{{.Logs}}</pre></blockquote></small>
		</div>
	</body>
</html>
`

var errHTMLTmpl = template.Must(template.New("").Funcs(tmplFuncs).Parse(errTmpl))

// ErrorHTML will render a list of code references
func ErrorHTML(w io.Writer, err error, logs string, css []byte) error {
	return errHTMLTmpl.Execute(w, struct {
		Title       string
		ErrorSimple string
		ErrorDetail string
		Logs        string
		Stylesheet  template.CSS
	}{
		Title:       "error occurred",
		ErrorSimple: err.Error(),
		ErrorDetail: fmt.Sprintf("%+v", errors.Cause(err)),
		Logs:        logs,
		Stylesheet:  template.CSS(string(css)),
	})
}
