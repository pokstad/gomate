package outline_test

import (
	"encoding/json"
	"testing"

	"github.com/pokstad/gomate"
	"github.com/pokstad/gomate/outline"
)

const expectedOutline = `[{"label":"outline","type":"package","start":78,"end":2745,"children":[{"label":"\"bytes\"","type":"import","start":105,"end":112},{"label":"\"fmt\"","type":"import","start":114,"end":119},{"label":"\"go/ast\"","type":"import","start":121,"end":129},{"label":"\"go/format\"","type":"import","start":131,"end":142},{"label":"\"go/parser\"","type":"import","start":144,"end":155},{"label":"\"go/token\"","type":"import","start":157,"end":167},{"label":"\"github.com/pokstad/gomate\"","type":"import","start":170,"end":197},{"label":"Declaration","type":"type","start":280,"end":587},{"label":"ParseFile","type":"function","start":652,"end":2475},{"label":"getReceiverType","type":"function","start":2477,"end":2745}]}]`

func TestParseFile(t *testing.T) {
	d, err := outline.ParseFile(gomate.Environment{CurrDoc: "outline.go"})
	if err != nil {
		t.Fatalf("unable to parse declarations: %s", err)
	}

	out, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("unable to marshal declarations into JSON: %s", err)
	}

	if string(out) != expectedOutline {
		t.Logf("expected: %s", expectedOutline)
		t.Logf("but got: %s", string(out))
		t.Fail()
	}
}
