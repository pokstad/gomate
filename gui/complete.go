package gui

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"text/template"

	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
)

// CompleteDialog shows a list of options to complete the current text. It is
// ported from the original Ruby version here:
// https://github.com/textmate/bundle-support.tmbundle/blob/d400d9d6a6234ccf3388d185c178b18a29078ada/Support/shared/lib/ui.rb#L128
type CompleteDialog struct {
	Choices         []Choice      // list of user selectable choices
	ExtraChars      string        // allowable characters besides alphanumeric
	CaseInsensitive bool          // ignore case when filtering suggestions
	StaticPrefix    string        // prefix used when filtering suggestions
	InitialFilter   string        // defaults to current word
	Images          []ChoiceImage // allowable images for choices
}

// dialogArgs will build a list of arguments for the dialog command
func (cd CompleteDialog) dialogArgs() ([]string, error) {
	args := []string{"popup"} // first arg indicates we want a popup dialog

	// choice suggestions
	choicesPlist, err := choicesToPlist(cd.Choices)
	if err != nil {
		return nil, errors.Wrap(err, "unable to render plist of popup choices")
	}
	args = append(args, "--suggestions", string(choicesPlist))

	if cd.ExtraChars != "" {
		args = append(args, "--additionalWordCharacters", cd.ExtraChars)
	}

	if cd.CaseInsensitive {
		args = append(args, "--caseInsensitive")
	}

	if cd.StaticPrefix != "" {
		args = append(args, "--staticPrefix", cd.StaticPrefix)
	}

	if cd.InitialFilter != "" {
		args = append(args, "--alreadyTyped", cd.InitialFilter)
	}

	return args, nil
}

// ChoiceImage is an image that can be used alongside a choice in the complete
// dialog
type ChoiceImage struct {
	Name string // unique name to reference image
	Path string // path to image file
}

// Choice is an item in the complete dialog a user may choose
type Choice struct {
	Display string // Display string for user
	Insert  string // string to insert after match/display
	Image   string // name of image to be shown alongside choice
	Match   string // typed text to filter on (defaults to display)
}

var popupPlistTmpl = template.Must(template.New("popup.plist").Parse(
	`({{ range .Choices}}{display = "{{ .Display }}"; insert = "{{ .Insert }}";},{{ end }})`,
))

func choicesToPlist(choices []Choice) ([]byte, error) {
	// TODO: should an error be returned if no choices are provided?

	b := new(bytes.Buffer)
	err := popupPlistTmpl.Execute(b, struct {
		Choices []Choice
	}{
		Choices: choices,
	})

	return b.Bytes(), err
}

func (cd CompleteDialog) registerImages(ctx context.Context, env gomate.Env) error {
	images := make([]string, len(cd.Images))
	for i := 0; i < len(cd.Images); i++ {
		images[i] = fmt.Sprintf(
			`%s = "%s";`,
			cd.Images[i].Name, cd.Images[i].Path,
		)
	}

	imagePlist := fmt.Sprintf(
		"{ %s }", strings.Join(images, " "),
	)

	cmd := exec.CommandContext(
		ctx,
		env.Dialog,
		"images",
		"--register", imagePlist,
	)

	stderrB := new(bytes.Buffer)
	cmd.Stderr = stderrB

	output, err := cmd.Output()
	if err != nil {
		log.Printf("textmate dialog image registration stderr: %s", stderrB.String())
		return errors.Wrap(err, "unable to register textmate dialog images")
	}

	log.Printf("textmate dialog image registration stdout: %s", output)

	return nil
}

// Show will display the completion dialog with all provided choices
func (cd CompleteDialog) Show(ctx context.Context, env gomate.Env) (string, error) {
	// if the choice dialog requires images, register them before showing dialog
	if len(cd.Images) > 0 {
		err := cd.registerImages(ctx, env)
		if err != nil {
			return "", errors.WithStack(err)
		}
	}

	args, err := cd.dialogArgs()
	if err != nil {
		return "", errors.Wrap(err, "unable to render dialog command args")
	}

	cmd := exec.CommandContext(
		ctx,
		env.Dialog,
		args...,
	)

	stderrB := new(bytes.Buffer)
	cmd.Stderr = stderrB

	out, err := cmd.Output()
	if err != nil {
		// log output to debug commands being constructed
		log.Printf("popup dialog stderr: %s", stderrB.String())
		return "", errors.Wrap(err, "unable to run textmate popup dialog")
	}

	return string(out), nil
}
