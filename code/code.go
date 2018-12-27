/*Package code provides code completion features.*/
package code

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/gui"
)

// ChoiceClass is an enumeration of all possible classes of identifiers
type ChoiceClass int

const (
	// ClassType indicates a type identifier
	ClassType ChoiceClass = iota + 1
	// ClassVar indicates a var identifier
	ClassVar
	// ClassFunc indicates a function identifier
	ClassFunc
	// ClassPackage indicates a package identifier
	ClassPackage
	// ClassConst indicates a constant identifier
	ClassConst
	// ClassPanic may indicate an error in the language server
	ClassPanic
)

// Choice is an option shown to the user to make a code completion selection
type Choice struct {
	Class   ChoiceClass // is it a type, var, func, or package?
	Display string      // what user is shown
	Insert  string      // what gets inserted into code
}

// Predictor is a type that can predict potential choices that make sense for
// the current cursor location
type Predictor interface {
	// Predict will provide a list of choices for the user to choose
	Predict(ctx context.Context, src io.Reader, env gomate.Env) ([]Choice, error)
}

// Completer provides code completion
type Completer struct {
	Predictor Predictor
}

func dialogChoices(choices []Choice) []gui.Choice {
	guiChoices := make([]gui.Choice, len(choices))

	for i := 0; i < len(choices); i++ {
		c := choices[i]
		guiChoices[i] = gui.Choice{
			Display: c.Display,
			Image:   classesGocode[c.Class],
			Match:   c.Insert,
		}
	}

	return guiChoices
}

// specific images are expected to be in the bundle support directory
func imagesFromEnv(env gomate.Env) ([]gui.ChoiceImage, error) {
	iconDir := filepath.Join(env.BundleDir, "icons")

	infos, err := ioutil.ReadDir(iconDir)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open bundle support icon dir")
	}

	var choices []gui.ChoiceImage
	for _, info := range infos {
		if info.IsDir() {
			continue // skip directories
		}

		comps := strings.Split(info.Name(), ".")
		if len(comps) < 1 {
			return nil, errors.Errorf("problem processing image path: %s", info.Name())
		}

		choices = append(choices, gui.ChoiceImage{
			Name: comps[0], // use base name of file as image name
			Path: filepath.Join(iconDir, info.Name()),
		})
	}

	return choices, nil
}

// Complete will call on the underlying predictor to provide a list of choices
// to the user. Once a selection has been made by the user, it will be inserted
// at the cursor by sending it to stdout.
func (c Completer) Complete(ctx context.Context, env gomate.Env, src io.Reader, dst io.Writer) error {
	choices, err := c.Predictor.Predict(ctx, src, env)
	if err != nil {
		return errors.Wrap(err, "unable to obtain predicted choices")
	}

	log.Printf("predictor choices: %#v", choices)

	images, err := imagesFromEnv(env)
	if err != nil {
		return errors.WithStack(err)
	}

	log.Printf("images: %#v", images)

	guiChoices := dialogChoices(choices)
	log.Printf("GUI dialog choices: %#v", guiChoices)

	dialog := gui.CompleteDialog{
		Choices:       guiChoices,
		InitialFilter: env.Cursor.Word,
		Images:        images,
	}

	choice, err := dialog.Show(ctx, env)
	if err != nil {
		return errors.Wrap(err, "unable to obtain selection from dialog")
	}

	log.Printf("selected choice: %s", choice)

	_, err = fmt.Fprintf(dst, choice)
	if err != nil {
		return errors.Wrap(err, "unable to write code completions to destination")
	}

	return nil
}
