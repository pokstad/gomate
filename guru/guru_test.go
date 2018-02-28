package guru_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/guru"
)

func TestParseReferrers(t *testing.T) {
	for _, tCase := range []struct {
		fPath  string
		line   uint
		column uint
		desc   string
	}{
		{
			fPath:  "guru.go",
			line:   19,
			column: 8,
			desc:   "package variable refRegex",
		},
	} {
		tCase := tCase // rescope iterator

		t.Run(tCase.desc, func(t *testing.T) {
			t.Logf("Parsing %s at line %d column %d",
				tCase.fPath, tCase.line, tCase.column,
			)

			refs, err := guru.ParseReferrers(
				context.Background(),
				gomate.Environment{
					CurrDoc:   "guru.go",
					CurrLine:  tCase.line,
					LineIndex: tCase.column,
					GoPath:    "/Users/paulokstad/go",
				},
			)
			if err != nil {
				t.Fatalf("cannot parse: %s", err)
			}

			buf := new(bytes.Buffer)

			err = gomate.RenderHTML(buf, nil, nil, refs)
			if err != nil {
				t.Fatalf("unable to render HTML: %s", err)
			}

			expected, err := ioutil.ReadFile("testdata/referrers.html")
			if err != nil {
				t.Fatalf("unable to open HTML fixture: %s", err)
			}

			if string(expected) != buf.String() {
				t.Logf("actual HTML: %s", buf.String())
				t.Fatalf("actual HTML does not match expected fixture")
			}
		})
	}
}
