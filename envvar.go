package gomate

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pokstad/offset/offset"
)

// Scope refers to the type of symbol under the cursor as determined by the
// textmate grammer
// See http://blog.macromates.com/2005/introduction-to-scopes/ for more info
type Scope string

const (
	// ScopeNone is when we don't receive a scope from textmate
	ScopeNone Scope = ""
)

// Cursor contains all information to determine where the UI cursor is
// currently pointing
type Cursor struct {
	Dir          string // directory of current doc (may be empty)
	Doc          string // filepath to document (may be empty)
	Line         int    // line number (starting at 1)
	Index        int    // index of line (starting at 1)
	Scope        Scope  // scope type specified by grammar
	Word         string // the word under the cursor (may be empty)
	SelectedText string // text currently highlighted by cursor (may be empty)
}

// RuneOffset attemps to calculate the rune offset for the cursor in the
// specified document. An error may occur if the document does not exist OR if
// the desired offset is out of bounds of the provided file.
func (c Cursor) RuneOffset() (int, error) {
	f, err := os.Open(c.Doc)
	if err != nil {
		return 0, errors.Wrap(err, "can't open file to determine offset")
	}
	defer f.Close() // nolint: errcheck

	offset, err := offset.CalcRune(f, c.Line, c.Index)
	if err != nil {
		return 0, errors.Wrap(err, "unable to calculate offset")
	}

	return offset, nil
}

// Drawer represents the drawer state
type Drawer struct {
	TopDir        string   // top level folder in project drawer
	Selected      []string // files and dirs currently selected
	SelectedFirst string   // first selected file (what determines first?)
}

// Env represents all required data derived from textmate/system environment
// variables
type Env struct {
	Cursor Cursor // where is current caret/cursor pointing?
	Drawer Drawer // current state of project drawer?

	// remaining dynamic env vars
	BundleDir   string // the support folder of the bundle that ran the item
	SoftTabs    bool   // are soft tabs being used?
	SupportPath string // common textmate support items
	TabSize     int    // size of soft tabs

	// Static Environment Variables
	GoPath string // GOPATH variable
	Dialog string // Path to Textmate dialog helper app
}

// GoBin returns the path to an executable in the $GOPATH/bin directory
func (e Env) GoBin(exe string) string {
	return filepath.Join(e.GoPath, "bin", exe)
}

// LoadEnvironment sources all environment variables or populates defaults
func LoadEnvironment() (env Env, err error) {
	return Env{
		Cursor: Cursor{
			Line:         parseInt(os.Getenv("TM_LINE_NUMBER")),
			Word:         os.Getenv("TM_CURRENT_WORD"),
			Dir:          os.Getenv("TM_DIRECTORY"),
			Doc:          os.Getenv("TM_FILEPATH"),
			Index:        parseInt(os.Getenv("TM_LINE_INDEX")) + 1,
			Scope:        Scope(os.Getenv("TM_SCOPE")),
			SelectedText: os.Getenv("TM_SELECTED_TEXT"),
		},

		Drawer: Drawer{
			TopDir:        os.Getenv("TM_PROJECT_DIRECTORY"),
			Selected:      strings.Split(os.Getenv("TM_SELECTED_FILES"), " "),
			SelectedFirst: os.Getenv("TM_SELECTED_FILE"),
		},

		BundleDir:   os.Getenv("TM_BUNDLE_SUPPORT"),
		SoftTabs:    os.Getenv("TM_SOFT_TABS") == "YES",
		SupportPath: os.Getenv("TM_SUPPORT_PATH"),
		TabSize:     parseInt(os.Getenv("TM_TAB_SIZE")),
		// gopath hierarchy: first trust textmate var, then try tradiational
		// GOPATH, finally assume $HOME/go as a last case effort
		GoPath: envOr("TM_GOPATH", envOr("GOPATH", filepath.Join(os.Getenv("HOME"), "go"))),
		Dialog: os.Getenv("DIALOG"),
	}, nil
}

func envOr(env, or string) string {
	e := os.Getenv(env)
	if e == "" {
		return or
	}
	return e
}

func parseInt(i string) int {
	d, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		return 0
	}
	return int(d)
}
