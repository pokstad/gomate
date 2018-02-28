package gomate

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Environment represents all textmate environment variables
type Environment struct {
	// Dynamic Environment Variables
	BundleDir     string   // the folder of the bundle that ran the item
	CurrLine      uint     // text contents of current line
	CurrWord      string   // text contents of current word
	CurrDir       string   // path to directory of current doc (may be empty)
	CurrDoc       string   // path to current doc (may be empty)
	LineIndex     uint     // index in the current line of the caret’s location
	ProjDir       string   // top level folder path in project drawer
	Scope         string   // scope selector the caret is inside
	SelectedFiles []string // files & dirs selected in the project drawer
	SelectedFile  string   // first selected file in project drawer
	SelectedText  string   // text currently selected
	SoftTabs      bool     // are soft tabs being used?
	SupportPath   string   // common textmate support items
	TabSize       uint     // size of soft tabs

	// Static Environment Variables
	GoPath string // GOPATH variable
}

// GoBin returns the path to the $GOPATH/bin directory
func (e Environment) GoBin() string {
	return filepath.Join(e.GoPath, "bin")
}

// LoadEnvironment sources all environment variables or populates defaults
func LoadEnvironment() (env Environment, err error) {
	defer func() {
		if e, ok := recover().(error); ok {
			err = e
		}
	}()

	return Environment{
		// Dynamic Vars
		BundleDir:     os.Getenv("TM_BUNDLE_SUPPORT"),
		CurrLine:      parseInt(os.Getenv("TM_CURRENT_LINE")),
		CurrWord:      os.Getenv("TM_CURRENT_WORD"),
		CurrDir:       os.Getenv("TM_DIRECTORY"),
		CurrDoc:       os.Getenv("TM_FILEPATH"),
		LineIndex:     parseInt(os.Getenv("TM_LINE_INDEX")),
		ProjDir:       os.Getenv("TM_PROJECT_DIRECTORY"),
		Scope:         os.Getenv("TM_SCOPE"),
		SelectedFiles: strings.Split(os.Getenv("TM_SELECTED_FILES"), " "),
		SelectedFile:  os.Getenv("TM_SELECTED_FILE"),
		SelectedText:  os.Getenv("TM_SELECTED_TEXT"),
		SoftTabs:      os.Getenv("TM_SOFT_TABS") == "YES",
		SupportPath:   os.Getenv("TM_SUPPORT_PATH"),
		TabSize:       parseInt(os.Getenv("TM_TAB_SIZE")),
		// Static Vars
		GoPath: func() string {
			p, err := filepath.Abs(envOr("TM_GOPATH", "~/go"))
			if err != nil {
				panic(PushE(err, "absolute path is elusive"))
			}
			return p
		}(),
	}, nil
}

func envOr(env, or string) string {
	e := os.Getenv(env)
	if e == "" {
		return or
	}
	return e
}

func parseInt(i string) uint {
	d, err := strconv.ParseInt(i, 10, 32)
	if err != nil {
		return 0
	}
	return uint(d)
}
