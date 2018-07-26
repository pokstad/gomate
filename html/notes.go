package html

import (
	"html/template"
	"io"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/notes"
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
	<h1>Notes: {{ .Env.Drawer.TopDir }}</h1>
	<ul>
		{{ range $pkg, $notes := .Notes }}
  		<li>{{ $pkg }}
        <ul>
  			{{ range $notes }}
          <li> <a href={{ print .Loc | safeURL }}>{{ .UID }}</a> : {{ .Body }}</li>
  			{{ end }}
        </ul>
  		</li>
		{{ end }}
	<ul>
</body>
</html>
`

var (
	notesTmpl = template.Must(template.New("").Funcs(tmplFuncs).Parse(notesHTML))
)

// RenderNotes will create a recursive outline of all godoc notes in the project
func RenderNotes(w io.Writer, env gomate.Env, pkgNotes map[string][]notes.Note, css []byte) error {
	return notesTmpl.Execute(w, struct {
		Stylesheet template.CSS
		Env        gomate.Env
		Notes      map[string][]notes.Note
	}{
		Stylesheet: template.CSS(string(css)),
		Env:        env,
		Notes:      pkgNotes,
	})
}
