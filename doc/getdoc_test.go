package doc_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/doc"
	"github.com/pokstad/gomate/test"
)

func TestGetDoc(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	getter, err := doc.ObtainGetter(gomate.Env{
		GoPath: os.Getenv("GOPATH"),
	})
	if err != nil {
		t.Fatalf("unable to obtain doc getter: %s", err)
	}

	for _, symCase := range []struct {
		desc     string
		cursor   gomate.Cursor
		expected doc.Symbol
	}{
		{
			desc: "struct type",
			cursor: gomate.Cursor{
				Doc:   test.MustGetAbsPath(t, "testdata/test.go"),
				Line:  27,
				Index: 7,
			},
			expected: doc.Symbol{
				Name:   "Meow",
				Import: "github.com/pokstad/gomate/doc/testdata",
				Pkg:    "main",
				Decl:   "type Meow struct {\n\tLoudness uint   // how loud is this meow?\n\tSounds   string // how would you spell the sound of this meow?\n}",
				Doc:    "Meow is an important test identifer for kitties\n",
				Pos: doc.Position{
					AbsPath: test.MustGetAbsPath(t, "testdata/test.go"),
					Line:    0x4,
					Column:  0x6,
				},
			},
		},
	} {
		t.Run(symCase.desc, func(t *testing.T) {
			symbol, err := getter.LookupSymbol(ctx, symCase.cursor)
			if err != nil {
				t.Fatalf("unable to obtain symbol: %s", err)
			}

			t.Logf("symbol returned by getter: %#v", symbol)
		})
	}

}
