package guru

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
)

// Guru is a wrapper around the guru command
type Guru struct {
	cmdPath string
}

// ObtainGuru will use the environment information to find the guru command
func ObtainGuru(env gomate.Env) (Guru, error) {
	g := Guru{
		cmdPath: env.GoBin("guru"),
	}
	return g, nil
}

func (g Guru) cmd(ctx context.Context, args ...string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	stdoutBuf := new(bytes.Buffer)
	stderrBuf := new(bytes.Buffer)

	cmd := exec.CommandContext(ctx, g.cmdPath, args...)

	cmd.Stderr = stderrBuf
	cmd.Stdout = stdoutBuf

	return cmd, stdoutBuf, stderrBuf
}

// captures absolute filepath, line number, column number, and source code clip
var refRegex = regexp.MustCompile(`^(.*?):(\d+)\.(\d+)-\d+\.\d+:\s+(.*?)$`)

// ParseReferrers uses guru to find code references to the specified symbol
// pointed at by the cursor. Results are returned with respect to the drawer.
func (g Guru) ParseReferrers(ctx context.Context, c gomate.Cursor, d gomate.Drawer) ([]gomate.CodeRef, error) {
	f, err := os.Open(c.Doc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open file")
	}

	offset, err := c.RuneOffset()
	if err != nil {
		return nil, errors.Wrap(err, "unable to calculate offset")
	}

	if err := f.Close(); err != nil {
		return nil, errors.Wrap(err, "unable to close file")
	}

	cmd, stdout, stderr := g.cmd(
		ctx,
		"referrers",
		fmt.Sprintf("%s:#%d", c.Doc, offset),
	)
	log.Printf("running command: %s", cmd.Args)

	err = cmd.Run()
	if err != nil {
		log.Printf("guru stderr: %s", string(stderr.Bytes()))

		return nil, errors.Wrap(err, "error from guru command")
	}

	var refs []gomate.CodeRef

	output := stdout.Bytes()

	log.Printf("guru output: %s", string(output))

	projDir := d.TopDir
	if projDir == "" {
		projDir = filepath.Dir(c.Doc)
	}

	for _, line := range strings.Split(string(output), "\n") {
		// check if empty line
		if len(strings.TrimSpace(line)) == 0 {
			// last line is empty
			break
		}

		m := refRegex.FindStringSubmatch(line)
		if len(m) < 4 {
			return nil, errors.Wrap(err, "unable to scan guru output")
		}

		var pErr error // int parse error

		atoi := func(s string) uint {
			var (
				i uint64
				e error
			)
			i, e = strconv.ParseUint(s, 10, 32)
			if e != nil {
				pErr = e
			}
			return uint(i)
		}

		relP, err := filepath.Rel(projDir, m[1])
		if err != nil {
			return nil, errors.Wrap(err, "unable to get relative path")
		}

		refs = append(refs, gomate.CodeRef{
			AbsPath:  m[1],
			Line:     atoi(m[2]),
			Column:   atoi(m[3]),
			RelPath:  relP, // TODO: trim project path from front
			Filename: path.Base(m[1]),
			Excerpt:  m[4],
		})

		if pErr != nil {
			return nil, errors.Wrap(pErr, "unable to convert string to int")
		}
	}

	return refs, nil
}
