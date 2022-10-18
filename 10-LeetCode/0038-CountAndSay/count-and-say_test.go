package countandsay

import "testing"

func Test_CountAndSay_S001(t *testing.T) {
	runCountAndSayTest(t, 1, "1") // base case
}

func Test_CountAndSay_S002(t *testing.T) {
	runCountAndSayTest(t, 2, "11") // one 1
}

func Test_CountAndSay_S003(t *testing.T) {
	runCountAndSayTest(t, 3, "21") // two 1s
}

func Test_CountAndSay_S004(t *testing.T) {
	runCountAndSayTest(t, 4, "1211") // one 2, one 1
}

func Test_CountAndSay_S005(t *testing.T) {
	runCountAndSayTest(t, 5, "111221") // one 1, one 2, 2 ones
}

// ----- helpers ----- //
func runCountAndSayTest(t *testing.T, input int, expected string) {
	t.Helper()

	var actual string = countAndSay(input)

	if actual != expected {
		t.Errorf("Input=%d Expected=%s Actual=%s", input, expected, actual)
	}
}
