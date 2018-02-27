package guru

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
)

func ParseRefernces() error {
	filepath := eVars.CurrDoc

	if !path.IsAbs(filepath) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		filepath = path.Join(wd, filepath)
	}

	refs, err := guruReferrers(context.TODO(), filepath, int(eVars.CurrLine), int(eVars.LineIndex))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to find referrers: %s", err)
		os.Exit(1)
	}

	if err := htmlTempl.Execute(os.Stdout, refs); err != nil {
		fmt.Fprintf(os.Stderr, "unable to render HTML: %s", err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout)

	return
}

type referrer struct {
	AbsPath  string
	Line     int
	Column   int
	RelPath  string
	Filename string
	Excerpt  string
}

// captures absolute filepath, line number, column number, and source code clip
var refRegex = regexp.MustCompile(`^(.*?):(\d+)\.(\d+)-\d+\.\d+:\s+(.*?)$`)

func guruReferrers(ctx context.Context, filepath string, line, col int) ([]referrer, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %s", filepath)
	}
	defer f.Close()

	offset, err := calcOffset(f, line, col)
	if err != nil {
		return nil, err
	}

	cmd := exec.CommandContext(
		ctx,
		"guru",
		"referrers",
		fmt.Sprintf("%s:#%d", path.Base(filepath), offset),
	)

	output, err := cmd.Output()
	if err != nil {
		log.Printf("error from guru command: %s", err)
		return nil, err
	}

	var refs []referrer

	log.Printf("Scan output: %s", string(output))

	for _, line := range strings.Split(string(output), "\n") {
		// check if empty line
		if len(strings.TrimSpace(line)) == 0 {
			// last line is empty
			break
		}

		m := refRegex.FindStringSubmatch(line)
		if len(m) < 4 {
			return nil, fmt.Errorf("unable to scan guru output: %s", line)
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

		refs = append(refs, referrer{
			AbsPath:  m[1],
			Line:     atoi(m[2]),
			Column:   atoi(m[3]),
			RelPath:  m[1], // TODO: trim project path from front
			Filename: path.Base(m[1]),
			Excerpt:  m[4],
		})

		if pErr != nil {
			return nil, fmt.Errorf("unable to parse number: %s", pErr)
		}
	}

	return refs, nil
}
