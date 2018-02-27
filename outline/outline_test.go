package outline_test

import (
	"encoding/json"
	"testing"

	"github.com/pokstad/gomate/outline"
)

const expectedOutline = `[{"label":"outline","type":"package","start":78,"end":2699,"children":[{"label":"\"bytes\"","type":"import","start":105,"end":112},{"label":"\"fmt\"","type":"import","start":114,"end":119},{"label":"\"go/ast\"","type":"import","start":121,"end":129},{"label":"\"go/format\"","type":"import","start":131,"end":142},{"label":"\"go/parser\"","type":"import","start":144,"end":155},{"label":"\"go/token\"","type":"import","start":157,"end":167},{"label":"Declaration","type":"type","start":250,"end":557},{"label":"ParseFile","type":"function","start":622,"end":2429},{"label":"getReceiverType","type":"function","start":2431,"end":2699}]}]`

func TestParseFile(t *testing.T) {
	d, err := outline.ParseFile("outline.go")
	if err != nil {
		t.Fatalf("unable to parse declarations: %s", err)
	}

	out, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("unable to marshal declarations into JSON: %s", err)
	}

	if string(out) != expectedOutline {
		t.Logf("expected: %s", string(out))
		t.Logf("but got: %s", expectedOutline)
		t.Fail()
	}
}
