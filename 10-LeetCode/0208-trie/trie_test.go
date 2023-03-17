package trie

import "testing"

func Test_One(t *testing.T) {

	var steps = []Step{
		*constructStep(),
		*insertStep("apple"),
		*searchStep("apple", true),
		*searchStep("app", false),
		*startsWithStep("app", true),
		*insertStep("app"),
		*searchStep("app", true),
	}

	runScenario(steps, t)
}

func Test_13(t *testing.T) {
	var steps = []Step{
		*constructStep(),
		*insertStep("app"),
		*insertStep("apple"),
		*insertStep("beer"),
		*insertStep("add"),
		*insertStep("jam"),
		*insertStep("rental"),
		*searchStep("apps", false),
		*searchStep("app", true),
		*searchStep("ad", false),
		*searchStep("applepie", false),
		*searchStep("rest", false),
		*searchStep("jan", false),
		*searchStep("rent", false),
		*searchStep("beer", true),
		*searchStep("jam", true),
		*startsWithStep("apps", false),
		*startsWithStep("app", true),
		*startsWithStep("ad", true),
		*startsWithStep("applepie", false),
		*startsWithStep("rest", false),
		*startsWithStep("jan", false),
		*startsWithStep("rent", true),
		*startsWithStep("beer", true),
		*startsWithStep("jam", true),
	}

	runScenario(steps, t)
}

func constructStep() *Step {
	return &(Step{"Trie", "", false})
}

func insertStep(word string) *Step {
	return &(Step{"insert", word, false})
}

func searchStep(word string, expected bool) *Step {
	return &(Step{"search", word, expected})
}

func startsWithStep(prefix string, expected bool) *Step {
	return &(Step{"startsWith", prefix, expected})
}

type Step struct {
	Command string
	Param   string
	Output  bool
}

func runScenario(steps []Step, t *testing.T) {
	t.Helper()

	var tr *Trie = nil

	for _, step := range steps {
		tr = executeStep(tr, &step, t)
	}
}

func executeStep(tr *Trie, step *Step, t *testing.T) *Trie {
	t.Helper()

	if step.Command == "Trie" {
		var obj = Constructor()
		return &obj
	} else if step.Command == "insert" {
		t.Logf("Insert(%s)", step.Param)
		tr.Insert(step.Param)
	} else if step.Command == "search" {
		print(tr, t, 0)
		t.Logf("Search(%s)", step.Param)
		var actual = tr.Search(step.Param)
		if actual != step.Output {
			t.Logf("Search(%s): Expected: %v Actual: %v", step.Param, step.Output, actual)
			t.Fail()
		}
	} else if step.Command == "startsWith" {
		t.Logf("StartsWith(%s)", step.Param)
		var actual = tr.StartsWith(step.Param)
		if actual != step.Output {
			t.Logf("StartsWith(%s): Expected: %v Actual: %v", step.Param, step.Output, actual)
			t.Fail()
		}
	} else {
		t.Errorf("Unknown step: %s", step.Command)
	}

	return tr
}

func print(tr *Trie, t *testing.T, depth int) {

	t.Logf("%*s (%v)", depth*4, string(tr.Letter), tr.IsWord)

	for i, _ := range tr.Next {
		print(tr.Next[i], t, depth+1)
	}
}
