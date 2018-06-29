package notes

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"

	"github.com/pokstad/gomate"
)

// Note is a godoc note found in Go source code comments
type Note struct {
	Tag  string
	UID  string
	Body string
	Loc  gomate.CodeRef
}

func (n Note) String() string {
	return fmt.Sprintf("%s: %s", n.UID, n.Body)
}

// AllNotes returns all notes for all packages in a directory
func AllNotes(target, base string) (map[string][]Note, error) {
	var (
		notes   = make(map[string][]Note)
		fileSet = token.NewFileSet()
	)

	pkgPaths := make(map[string]map[string]*ast.Package)

	// walk entire target contents to parse all contents
	err := filepath.Walk(target,
		func(p string, info os.FileInfo, err error) error {
			switch {

			case err != nil:
				log.Printf("unable to walk path %s due to :%s", p, err)
				return err

			case !info.IsDir():
				return nil // skip files

			case info.Name() == "testdata":
				return filepath.SkipDir // skip fixtures

			case info.Name() == "vendor":
				return filepath.SkipDir // skip vendored deps

			}

			pkgs, err := parser.ParseDir(fileSet, p, nil, parser.ParseComments)
			if err != nil && pkgs == nil {
				return gomate.PushE(err, "unable to parse dir")
			}

			pkgPaths[p] = pkgs

			return nil
		},
	)
	if err != nil {
		return nil, gomate.PushE(err, "unable to walk directory")
	}

	for _, pkgs := range pkgPaths {
		for pkgName, pkgAST := range pkgs {
			d := doc.New(pkgAST, pkgName, 0)
			for m, ns := range d.Notes {
				for _, n := range ns {
					t := fileSet.Position(n.Pos)

					relPath, err := filepath.Rel(base, t.Filename)
					if err != nil {
						return nil, gomate.PushE(
							err,
							"can't determine relative path",
						)
					}

					absPath, err := filepath.Abs(t.Filename)
					if err != nil {
						return nil, gomate.PushE(err, "can't get absolute path")
					}

					notes[m] = append(notes[m], Note{
						Tag:  m,
						UID:  n.UID,
						Body: n.Body,
						Loc: gomate.CodeRef{
							AbsPath:  absPath,
							Column:   uint(t.Column),
							Excerpt:  fmt.Sprintf("%s(%s): %s", m, n.UID, n.Body),
							Filename: filepath.Base(t.Filename),
							Line:     uint(t.Line),
							RelPath:  relPath,
						},
					})
				}
			}
		}
	}

	return notes, nil
}

//func sourceForPos(start, end token.Position)

//func RecursiveNotes()
