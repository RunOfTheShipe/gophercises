package stringcompression

import (
	"fmt"
	"testing"
)

func getLengthOfOptimalCompression(s string, k int) int {

	var solutions = make(map[string]string)
	return getLengthOfOptimalCompressionRecurse(s, k, solutions)
}

func getLengthOfOptimalCompressionRecurse(s string, k int, solutions map[string]string) int {
	var inputLength = len(s)

	if inputLength == 0 {
		return 0
	} else if inputLength-k == 1 {
		return 1
	} else {
		var localCompress = runLengthCompressionWithLookup(s, solutions)
		var localSolution = len(localCompress)

		if localCompress == s {
			// nothing could be compressed, so the best solution
			// is to just remove the first k characters
			return inputLength - k
		}

		if k == 0 {
			// can't remove anything, so best option is to compress it as it already is
			return localSolution
		} else {
			// can remove at least one character, so iterate and see if any
			// solution with removing a single character gets us anything
			// shorter than what we already have
			for i := 0; i < inputLength; i++ {
				var tempSolution int = -1

				if i == 0 {
					tempSolution = getLengthOfOptimalCompressionRecurse(s[1:], k-1, solutions)
				} else if i == inputLength-1 {
					tempSolution = getLengthOfOptimalCompressionRecurse(s[:inputLength-1], k-1, solutions)
				} else {
					// have to build a string based on two splices
					var temp = fmt.Sprintf("%s%s", s[:i], s[i+1:])
					tempSolution = getLengthOfOptimalCompressionRecurse(temp, k-1, solutions)
				}

				if tempSolution < localSolution {
					localSolution = tempSolution
				}
			}
			return localSolution
		}
	}
}

var misses int = 0
var calls int = 0

func runLengthCompressionWithLookup(s string, lookup map[string]string) string {
	calls++
	solution, ok := lookup[s]
	if !ok {
		misses++
		solution = runLengthCompression(s)
		lookup[s] = solution
	}
	return solution
}

func runLengthCompression(s string) string {
	var compressed string = ""

	var lastRune rune = rune(s[0])
	var lastRuneCount int = 1

	for _, currentRune := range s[1:] {
		if lastRune == currentRune {
			lastRuneCount++
		} else {
			// new rune detected - updated the compressed string
			if lastRuneCount == 1 {
				compressed = fmt.Sprintf("%s%c", compressed, lastRune)
			} else {
				compressed = fmt.Sprintf("%s%c%d", compressed, lastRune, lastRuneCount)
			}

			// restart rune counting with the current
			lastRune = currentRune
			lastRuneCount = 1
		}
	}

	// append the last rune/count
	if lastRuneCount == 1 {
		compressed = fmt.Sprintf("%s%c", compressed, lastRune)
	} else {
		compressed = fmt.Sprintf("%s%c%d", compressed, lastRune, lastRuneCount)
	}

	return compressed
}

// ----- testing ----- //
func Test_S001(t *testing.T) {
	testGetLength(t, "aaabcccd", 2, 4)
}

func Test_S002(t *testing.T) {
	testGetLength(t, "aabbaa", 2, 2)
}

func Test_S003(t *testing.T) {
	testGetLength(t, "aaaaaaaaaaa", 0, 3)
}

func Test_S004(t *testing.T) {
	testGetLength(t, "aaaaaaaaaaabc", 3, 3)
}

func Test_S005(t *testing.T) {
	testGetLength(t, "a", 1, 0)
}

func Test_S006(t *testing.T) {
	testGetLength(t, "abc", 2, 1)
}

func Test_S007(t *testing.T) {
	testGetLength(t, "abcdefghijklmnopqrstuvwxyz", 16, 10)
}

func Test_S008(t *testing.T) {
	testGetLength(t, "aaccbcacbaab", 11, 1)
}

func Test_S009(t *testing.T) {
	testGetLength(t, "cdacddcaaacbc", 7, 4)
}

func Test_RunCompression_S001(t *testing.T) {
	testCompression(t, "aaabcccd", "a3bc3d")
}

func Test_RunCompression_S002(t *testing.T) {
	testCompression(t, "aabbaa", "a2b2a2")
}

func Test_RunCompression_S003(t *testing.T) {
	testCompression(t, "aaaaaaaaaaa", "a11")
}

func Test_RunCompression_S004(t *testing.T) {
	testCompression(t, "aaccbcacbaab", "a2c2bcacba2b")
}

func Test_RunCompression_S005(t *testing.T) {
	testCompression(t, "cdacddcaaacbc", "cdacd2ca3cbc")
}

// ----- helpers ----- //
func testGetLength(t *testing.T, input string, k int, expected int) {
	misses = 0
	calls = 0

	var actual = getLengthOfOptimalCompression(input, k)
	if actual != expected {
		t.Errorf("Input=%s K=%d Expected=%d Actual=%d", input, k, expected, actual)
	}
	t.Logf("Calls=%d Misses=%d", calls, misses)
}

func testCompression(t *testing.T, input string, expected string) {
	var actual = runLengthCompression(input)
	if actual != expected {
		t.Errorf("Input=%s Expected=%s Actual=%s", input, expected, actual)
	}
}
