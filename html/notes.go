package html

import (
	"html/template"
	"io"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/outline"
)

const notesHTML = `
<!doctype html>
<html lang="en">
<head>
  <title>Outline for {{ .Env.Cursor.Doc }}</title>
  <style type="text/css">
	{{ .Stylesheet }}
  </style>
</head>
<body class="remarkdown">
	<h1>Outline: {{ .Env.Cursor.Doc }}</h1>
	<ul>
		{{ define "declaration" }}
			<li>{{ .Label }}
				{{ if gt (len .Children) 0}}
					<ul>
						{{ range .Children }}
						{{ template "declaration" . }}
						{{ end }}
					</ul>
				{{ end }}
			</li>
		{{ end }}
		{{ range .Declarations }}
			{{ template "declaration" . }}
		{{ end }}
	<ul>
</body>
</html>
`

var (
	notesTmpl = template.Must(template.New("").Parse(notesHTML))
)

// OutlineHTML will create a recursive outline of all symbols in the current
// source code file
func RenderNotes(w io.Writer, env gomate.Env, decs []outline.Decl, css []byte) error {
	return outlineTmpl.Execute(w, struct {
		Stylesheet   template.CSS
		Env          gomate.Env
		Declarations []outline.Decl
	}{
		Stylesheet:   template.CSS(string(css)),
		Env:          env,
		Declarations: decs,
	})
}
