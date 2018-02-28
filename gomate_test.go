package gomate_test

import (
	"bytes"
	"testing"

	"github.com/pokstad/gomate"
)

func TestRenderHTML(t *testing.T) {
	buf := new(bytes.Buffer)

	err := gomate.RenderHTML(buf, nil, nil, []gomate.CodeRef{
		{
			AbsPath:  "/test.txt",
			Line:     1,
			Column:   5,
			RelPath:  "test.txt",
			Filename: "test.txt",
			Excerpt:  "this is a test",
		},
	})
	if err != nil {
		t.Fatalf("unable to render HTML: %s", err)
	}

	if buf.String() != expectedHTML {
		t.Fatalf("did not expect HTML output: %s", buf.String())
	}

}

const expectedHTML = `
<html>
<head>
</head>
<body>

<div>
<h2> errors </h2>
<ul>

</ul>
</div>

<div>
<h2> infos </h2>
<ul>

<li><a href="txmt://open?line=1&amp;column=5&amp;url=file:///test.txt&amp;title=test.txt">test.txt:1:5:</a> this is a test</li>

</ul>
</div>

<div>
<h2> warnings </h2>
<ul>

</ul>
</div>

</body></html>
`
