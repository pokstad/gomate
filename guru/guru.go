package guru

import (
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

	"github.com/pokstad/gomate"
)

// captures absolute filepath, line number, column number, and source code clip
var refRegex = regexp.MustCompile(`^(.*?):(\d+)\.(\d+)-\d+\.\d+:\s+(.*?)$`)

// ParseReferrers uses guru to find code references to the specified symbol
// pointed at in the file
func ParseReferrers(ctx context.Context, env gomate.Env) ([]gomate.CodeRef, error) {
	fPath, err := filepath.Abs(env.CurrDoc)
	if err != nil {
		return nil, gomate.PushE(err, "unable to get working directory")
	}

	f, err := os.Open(fPath)
	if err != nil {
		return nil, gomate.PushE(err, "unable to open file")
	}

	offset, err := gomate.CalcOffset(f, env.CurrLine, env.LineIndex)
	if err != nil {
		return nil, gomate.PushE(err, "unable to calculate offset")
	}

	if err := f.Close(); err != nil {
		return nil, gomate.PushE(err, "unable to close file")
	}

	args := []string{
		path.Join(env.GoBin(), "guru"),
		"referrers",
		fmt.Sprintf("%s:#%d", path.Base(fPath), offset),
	}

	log.Printf("running command: %s", args)

	cmd := exec.CommandContext(
		ctx,
		args[0],
		args[1:]...,
	)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("guru failed: %s", string(output))
		return nil, gomate.PushE(err, "error from guru command")
	}

	var refs []gomate.CodeRef

	log.Printf("guru output: %s", string(output))

	for _, line := range strings.Split(string(output), "\n") {
		// check if empty line
		if len(strings.TrimSpace(line)) == 0 {
			// last line is empty
			break
		}

		m := refRegex.FindStringSubmatch(line)
		if len(m) < 4 {
			return nil, gomate.PushE(err, "unable to scan guru output")
		}

		var pErr error // int parse error

		atoi := func(s string) int {
			var (
				i int64
				e error
			)
			i, e = strconv.ParseInt(s, 10, 32)
			if e != nil {
				pErr = e
			}
			return int(i)
		}

		relP, err := filepath.Rel(env.CurrDir, m[1])
		if err != nil {
			return nil, gomate.PushE(err, "unable to get relative path")
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
			return nil, gomate.PushE(pErr, "unable to convert string to int")
		}
	}

	return refs, nil
}
