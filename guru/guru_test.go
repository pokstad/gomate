package guru_test

import (
	"context"
	"os"
	"path/filepath"
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
		expect []gomate.CodeRef
	}{
		{
			fPath:  "guru.go",
			line:   19,
			column: 8,
			desc:   "package variable refRegex",
			expect: []gomate.CodeRef{
				{
					AbsPath:  mustGetAbsPath("guru.go"),
					Line:     19,
					Column:   5,
					RelPath:  "guru.go",
					Filename: "guru.go",
					Excerpt:  "references to var refRegex *regexp.Regexp",
				},
				{
					AbsPath:  mustGetAbsPath("guru.go"),
					Line:     74,
					Column:   8,
					RelPath:  "guru.go",
					Filename: "guru.go",
					Excerpt:  "m := refRegex.FindStringSubmatch(line)",
				},
			},
		},
	} {
		tCase := tCase // rescope iterator

		t.Run(tCase.desc, func(t *testing.T) {
			t.Logf("Parsing %s at line %d column %d",
				tCase.fPath, tCase.line, tCase.column,
			)

			refs, err := guru.ParseReferrers(
				context.Background(),
				gomate.Env{
					CurrDoc:   "guru.go",
					CurrLine:  tCase.line,
					LineIndex: tCase.column,
					GoPath:    os.Getenv("GOPATH"),
					CurrDir: func() string {
						rp, err := os.Getwd()
						if err != nil {
							panic(err)
						}
						return rp
					}(),
				},
			)
			if err != nil {
				t.Fatalf("cannot parse: %s", err)
			}

			for i, e := range tCase.expect {
				if refs[i] != e {
					t.Fatalf("mismatch: %#v vs %#v", e, refs[i])
				}
			}
		})
	}
}

func mustGetAbsPath(relP string) string {
	p, err := filepath.Abs(relP)
	if err != nil {
		panic(err)
	}
	return p
}
