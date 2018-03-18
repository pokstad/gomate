package doc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go/doc"
	"log"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/pokstad/gomate"
)

// Symbol represents documentation about a Go source code symbol
type Symbol struct {
	Name   string   `json:"name"`
	Import string   `json:"import"`
	Pkg    string   `json:"pkg"`
	Decl   string   `json:"decl"`
	Doc    string   `json:"doc"`
	Pos    Position `json:"pos"`
}

// Position points to where this symbol is
type Position struct {
	AbsPath string
	Line    uint
	Column  uint
}

func (s Symbol) HTML() string {
	b := new(bytes.Buffer)
	doc.ToHTML(b, s.Doc, nil)
	return b.String()
}

func (s Symbol) CodeRef(baseDir string) gomate.CodeRef {
	if baseDir == "" {
		baseDir = filepath.Dir(s.Pos.AbsPath)
	}

	rp, err := filepath.Rel(baseDir, s.Pos.AbsPath)
	if err != nil {
		log.Printf("unable to obtain relative path: %s", err)
		rp = s.Pos.AbsPath
	}

	return gomate.CodeRef{
		AbsPath:  s.Pos.AbsPath,
		Column:   s.Pos.Column,
		Excerpt:  fmt.Sprintf("%s.%s", s.Pkg, s.Name),
		Filename: filepath.Base(s.Pos.AbsPath),
		Line:     s.Pos.Line,
		RelPath:  rp,
	}
}

var posRegex = regexp.MustCompile(`^"(.+?):(\d+):(\d+)"$`)

func (p *Position) UnmarshalJSON(data []byte) error {
	m := posRegex.FindStringSubmatch(string(data))
	if len(m) != 4 {
		return gomate.E("unexpected number of regexp matches for position: " + strconv.Itoa(len(m)))
	}

	line, err := strconv.Atoi(m[2])
	if err != nil {
		return gomate.PushE(err, "unable to parse line number: "+err.Error())
	}

	index, err := strconv.Atoi(m[3])
	if err != nil {
		return gomate.PushE(err, "unable to parse index number: "+err.Error())
	}

	*p = Position{
		AbsPath: m[1],
		Line:    uint(line),
		Column:  uint(index),
	}

	return nil
}

type Getter struct {
	cmdPath string
}

// ObtainGetter will attempt to return a reference to the environment doc
// getter.
func ObtainGetter(env gomate.Env) (Getter, error) {
	return Getter{cmdPath: env.GoBin("gogetdoc")}, nil
}

// LookupSymbol attempts to look up a symbol's doc via file and offset
func (g Getter) LookupSymbol(ctx context.Context, c gomate.Cursor) (Symbol, error) {
	offset, err := c.RuneOffset()
	if err != nil {
		return Symbol{}, gomate.PushE(err, "can't determine offset")
	}

	cmd := exec.CommandContext(
		ctx,
		g.cmdPath,
		"-pos",
		fmt.Sprintf("%s:#%d",
			c.Doc,
			offset,
		),
		"-json",
	)

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	log.Printf("running gogetdoc command: %s", cmd.Args)

	err = cmd.Run()
	if err != nil {
		log.Printf("gogetdoc stderr: %s", stderr.String())
		return Symbol{}, gomate.PushE(err, "unable to run gogetdoc")
	}

	var d Symbol

	log.Printf("decoding json from stdout: %s", stdout.String())

	if err := json.NewDecoder(stdout).Decode(&d); err != nil {
		return Symbol{}, gomate.PushE(err, "unable to decode gogetdoc JSON output")
	}

	return d, nil
}
