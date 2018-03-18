package guru_test

import (
	"context"
	"os"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/guru"
	"github.com/pokstad/gomate/test"
)

func TestParseReferrers(t *testing.T) {
	g, err := guru.ObtainGuru(gomate.Env{
		GoPath: os.Getenv("GOPATH"),
	})
	if err != nil {
		t.Fatalf("cannot obtain guru reference")
	}

	for _, tCase := range []struct {
		desc   string
		cursor gomate.Cursor
		drawer gomate.Drawer
		expect []gomate.CodeRef
	}{
		{
			desc: "type declaration for sample",
			cursor: gomate.Cursor{
				Doc:   test.MustGetAbsPath(t, "testdata/test.go"),
				Line:  3,
				Index: 9,
			},
			drawer: gomate.Drawer{
				TopDir: func() string {
					wd, err := os.Getwd()
					if err != nil {
						panic(err)
					}
					return wd
				}(),
			},
			expect: []gomate.CodeRef{
				{
					AbsPath:  test.MustGetAbsPath(t, "testdata/test.go"),
					Line:     3,
					Column:   6,
					RelPath:  "testdata/test.go",
					Filename: "test.go",
					Excerpt:  "references to type sample struct{}",
				},
				{
					AbsPath:  test.MustGetAbsPath(t, "testdata/test.go"),
					Line:     5,
					Column:   9,
					RelPath:  "testdata/test.go",
					Filename: "test.go",
					Excerpt:  "func (s sample) Do() {}",
				},
				{
					AbsPath:  test.MustGetAbsPath(t, "testdata/test.go"),
					Line:     8,
					Column:   2,
					RelPath:  "testdata/test.go",
					Filename: "test.go",
					Excerpt:  "sample{}.Do()",
				},
			},
		},
		{
			desc: "func main declaration with no drawer topdir",
			cursor: gomate.Cursor{
				Doc:   test.MustGetAbsPath(t, "testdata/test.go"),
				Line:  7,
				Index: 6,
			},
			drawer: gomate.Drawer{
			// intentionally empty to simulate a single file (no project)
			},
			expect: []gomate.CodeRef{
				{
					AbsPath:  test.MustGetAbsPath(t, "testdata/test.go"),
					Line:     7,
					Column:   6,
					RelPath:  "test.go",
					Filename: "test.go",
					Excerpt:  "references to func main()",
				},
			},
		},
	} {
		t.Run(tCase.desc, func(t *testing.T) {
			t.Logf("Parsing %+v with drawer %+v",
				tCase.cursor, tCase.drawer,
			)

			refs, err := g.ParseReferrers(
				context.Background(),
				tCase.cursor,
				tCase.drawer,
			)
			if err != nil {
				t.Fatalf("cannot parse: %s", err)
			}

			t.Logf("refs: %#v", refs)

			if len(tCase.expect) != len(refs) {
				t.Fatalf("expected %d results but got %d",
					len(tCase.expect), len(refs),
				)
			}

			for i, e := range tCase.expect {
				if refs[i] != e {
					t.Logf("expected: %+v", e)
					t.Fatalf("mismatch: %#v vs %#v", e, refs[i])
				}
			}
		})
	}
}
