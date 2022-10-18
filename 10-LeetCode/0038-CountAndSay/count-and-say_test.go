package countandsay

import "testing"

func countAndSay(n int) string {
	return ""
}

// ----- testing ----- //
func Test_CountAndSay_S001(t *testing.T) {
	runCountAndSayTest(t, 1, "1")
}

func Test_CountAndSay_S002(t *testing.T) {
	runCountAndSayTest(t, 4, "1121")
}

// ----- helpers ----- //
func runCountAndSayTest(t *testing.T, input int, expected string) {
	t.Helper()

	var actual string = countAndSay(input)

	if actual != expected {
		t.Errorf("Input=%d Expected=%s Actual=%s", input, expected, actual)
	}
}
