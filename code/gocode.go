package code

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"

	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
	"github.com/pokstad/offset/offset"
)

// GoCode is the code predictor wrapper for https://github.com/stamblerre/gocode
type GoCode struct {
}

// Predict will return a list of choices for the position at the file
func (gc GoCode) Predict(ctx context.Context, src io.Reader, env gomate.Env) ([]Choice, error) {
	srcBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read source")
	}

	log.Printf("src: %s", string(srcBytes))

	srcRdr := bytes.NewReader(srcBytes)

	off, err := offset.CalcRune(srcRdr, env.Cursor.Line, env.Cursor.Index)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to calculate offset in source for line %d and column %d", env.Cursor.Line, env.Cursor.Index)
	}

	gcArgs := []string{
		"-f", "json", // JSON format output
		"-source", // import source code rather than compiled package
		"autocomplete",
		env.Cursor.Doc, strconv.FormatInt(int64(off), 10), // where cursor is
	}

	log.Printf("gocode args: %s", gcArgs)

	cmd := exec.CommandContext(
		ctx,
		"gocode",
		gcArgs...,
	)

	srcRdr.Reset(srcBytes)
	cmd.Stdin = srcRdr

	stderrB := new(bytes.Buffer)
	stdoutB := new(bytes.Buffer)
	cmd.Stderr = stderrB
	cmd.Stdout = stdoutB

	err = cmd.Run()
	log.Printf("gocode stdout: %s", stdoutB.String())
	log.Printf("gocode stderr: %s", stderrB.String())
	if err != nil {

		return nil, errors.Wrap(err, "gocode command failed")
	}

	choices, err := ParseGocode(stdoutB.Bytes())
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse gocode output")
	}

	log.Printf("choices: %#v", choices)

	return choices, nil
}

type gocodeChoice struct {
	Class   string
	Type    string
	Name    string
	Package string
}

var gocodeClasses = map[string]ChoiceClass{
	"type":    ClassType,
	"var":     ClassVar,
	"func":    ClassFunc,
	"package": ClassPackage,
	"const":   ClassConst,
	"PANIC":   ClassPanic,
}

var classesGocode = map[ChoiceClass]string{
	ClassType:    "type",
	ClassVar:     "var",
	ClassFunc:    "func",
	ClassPackage: "package",
	ClassConst:   "const",
	ClassPanic:   "PANIC",
}

func choiceFromGocode(gc gocodeChoice) (Choice, error) {
	c := Choice{
		Display: fmt.Sprintf("%s - %s", gc.Name, gc.Type),
		Insert:  gc.Name,
	}

	c.Class = gocodeClasses[gc.Class]
	if c.Class == 0 {
		return Choice{}, errors.Errorf("gocode result contains invalid class: %s", gc.Class)
	}

	return c, nil
}

// ParseGocode will parse the output of gocode for code completion choices
func ParseGocode(output []byte) ([]Choice, error) {
	array := [2]json.RawMessage{}

	err := json.Unmarshal(output, &array)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode JSON array")
	}

	results := []gocodeChoice{}

	err = json.Unmarshal(array[1], &results)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode JSON gocode results")
	}

	log.Printf("choices: %#v", results)

	choices := make([]Choice, len(results))
	for i := 0; i < len(choices); i++ {
		choices[i], err = choiceFromGocode(results[i])
		if err != nil {
			return nil, errors.Wrap(err, "unable to convert gocode choice")
		}
	}

	return choices, nil
}
