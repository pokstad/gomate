package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var htmlTempl = template.Must(template.New("").Parse(`
<html>
<head>
</head>
<body>
  <div>
    <ul>
      {{ range . }}
      <li><a href="txmt://open?line={{ .Line }}&amp;column={{ .Column }}&amp;url=file://{{ .AbsPath }}" title="{{ .RelPath }}">{{ .Filename }}:{{ .Line }}:{{ .Column }}:</a> {{ .Excerpt }}</li>
      {{ end }}
    </ul>
  </div>
</body></html>
`))

func calcOffset(r io.Reader, line, col int) (int, error) {
	var (
		s  = bufio.NewScanner(r)
		bc = 0
	)

	for i := 1; s.Scan() && i < line; i++ {
		if err := s.Err(); err != nil {
			return 0, fmt.Errorf("unable to scan file for offset: %s", err)
		}

		bc += len(s.Text()) + 1 // add newline to text length
	}

	return bc + col, nil
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

func main() {
	var (
		filepath = flag.String("path", "", "path to source code file")
		line     = flag.Int("line", 0, "line number")
		col      = flag.Int("column", 0, "column position")
	)

	flag.Parse()

	if !path.IsAbs(*filepath) {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		*filepath = path.Join(wd, *filepath)
	}

	refs, err := guruReferrers(context.TODO(), *filepath, *line, *col)
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
