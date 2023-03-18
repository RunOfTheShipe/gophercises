package browserhistory

import (
	"strconv"
	"testing"
)

func TestOne(t *testing.T) {
	testScenario(t,
		Step{"BrowserHistory", "leetcode.com", ""},
		Step{"visit", "google.com", ""},
		Step{"visit", "facebook.com", ""},
		Step{"visit", "youtube.com", ""},
		Step{"back", "1", "facebook.com"},
		Step{"back", "1", "google.com"},
		Step{"forward", "1", "facebook.com"},
		Step{"visit", "linkedin.com", ""},
		Step{"forward", "2", "linkedin.com"},
		Step{"back", "2", "google.com"},
		Step{"back", "7", "leetcode.com"})
}

type Step struct {
	Action         string
	Param          string
	ExpectedOutput string
}

func testScenario(t *testing.T, steps ...Step) {
	var bh *BrowserHistory = nil

	for _, step := range steps {
		bh = executeStep(bh, step, t)
	}
}

func executeStep(bh *BrowserHistory, step Step, t *testing.T) *BrowserHistory {
	t.Helper()

	if step.Action == "BrowserHistory" {
		t.Logf("Create(%s)", step.Param)
		var b = Constructor(step.Param)
		return &b
	} else if step.Action == "visit" {
		t.Logf("Visit(%s)", step.Param)
		bh.Visit(step.Param)
	} else {
		t.Logf("%s(%s)", step.Action, step.Param)
		var steps, err = strconv.Atoi(step.Param)
		if err != nil {
			t.Errorf("%v", err)
		}

		var actual string = ""
		if step.Action == "back" {
			actual = bh.Back(steps)
		} else {
			actual = bh.Forward(steps)
		}

		if actual != step.ExpectedOutput {
			t.Logf("%s(%s): Expected: %s Actual: %s", step.Action, step.Param, step.ExpectedOutput, actual)
			t.Fail()
		}

	}

	return bh
}
