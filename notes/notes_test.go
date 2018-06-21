package notes_test

import (
	"testing"

	"github.com/pokstad/gomate/notes"
)

func TestAllNotes(t *testing.T) {
	allNotes, err := notes.AllNotes("../", "")
	if err != nil {
		t.Fatalf("unable to parse all notes: %s", err)
	}

	t.Logf("%#v", allNotes)
}
