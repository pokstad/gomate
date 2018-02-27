/*Package outline is adapted from https://github.com/lukehoban/go-outline */
package outline

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
)

// Declaration represents a top level declaration of the source code file
type Declaration struct {
	Label        string        `json:"label"`
	Type         string        `json:"type"`
	ReceiverType string        `json:"receiverType,omitempty"`
	Start        token.Pos     `json:"start"`
	End          token.Pos     `json:"end"`
	Children     []Declaration `json:"children,omitempty"`
}

// ParseFile will parse a Go source code file for declarations
func ParseFile(fPath string) ([]Declaration, error) {
	fset := token.NewFileSet()
	parserMode := parser.ParseComments

	fileAst, err := parser.ParseFile(fset, fPath, nil, parserMode)
	if err != nil {
		return nil, fmt.Errorf("unable to parse declarations: %s", err)
	}

	declarations := []Declaration{}

	for _, decl := range fileAst.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			receiverType, err := getReceiverType(fset, decl)
			if err != nil {
				return nil, fmt.Errorf("Failed to parse receiver type: %v", err)
			}
			declarations = append(declarations, Declaration{
				decl.Name.String(),
				"function",
				receiverType,
				decl.Pos(),
				decl.End(),
				[]Declaration{},
			})
		case *ast.GenDecl:
			for _, spec := range decl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					declarations = append(declarations, Declaration{
						spec.Path.Value,
						"import",
						"",
						spec.Pos(),
						spec.End(),
						[]Declaration{},
					})
				case *ast.TypeSpec:
					//TODO: Members if it's a struct or interface type?
					declarations = append(declarations, Declaration{
						spec.Name.String(),
						"type",
						"",
						spec.Pos(),
						spec.End(),
						[]Declaration{},
					})
				case *ast.ValueSpec:
					for _, id := range spec.Names {
						declarations = append(declarations, Declaration{
							id.Name,
							"variable",
							"",
							id.Pos(),
							id.End(),
							[]Declaration{},
						})
					}
				default:
					return nil, fmt.Errorf("Unknown token type: %s", decl.Tok)
				}
			}
		default:
			return nil, fmt.Errorf("Unknown declaration @", decl.Pos())
		}
	}

	return []Declaration{
		Declaration{
			fileAst.Name.String(),
			"package",
			"",
			fileAst.Pos(),
			fileAst.End(),
			declarations,
		},
	}, nil
}

func getReceiverType(fset *token.FileSet, decl *ast.FuncDecl) (string, error) {
	if decl.Recv == nil {
		return "", nil
	}

	buf := &bytes.Buffer{}
	if err := format.Node(buf, fset, decl.Recv.List[0].Type); err != nil {
		return "", err
	}

	return buf.String(), nil
}
