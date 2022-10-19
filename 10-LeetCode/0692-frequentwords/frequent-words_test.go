package frequentwords

import "testing"

func Test_FrequentWords_S001(t *testing.T) {
	testFrequentWords(t,
		[]string{"i", "love", "leetcode", "i", "love", "coding"}, 2,
		[]string{"i", "love"})
}

func Test_FrequentWords_S002(t *testing.T) {
	testFrequentWords(t,
		[]string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}, 4,
		[]string{"the", "is", "sunny", "day"})
}

// k is bigger than the number of unique words
func Test_FrequentWords_S011(t *testing.T) {
	testFrequentWords(t,
		[]string{"the", "the", "the", "the", "the", "the", "the", "the", "the", "the"}, 4,
		[]string{"the"})
}

// all the words are unique
func Test_FrequentWords_S012(t *testing.T) {
	testFrequentWords(t,
		[]string{"z", "y", "x", "w", "v", "u", "t", "s", "r", "q"}, 4,
		[]string{"q", "r", "s", "t"})
}

// ----- helpers ----- //

func testFrequentWords(t *testing.T, words []string, k int, expected []string) {
	t.Helper()

	var wordCopy []string = make([]string, len(words))
	copy(wordCopy, words)

	var actual []string = topKFrequent(wordCopy, k)

	var areTheSame bool = len(actual) == len(expected)
	if areTheSame {
		for i := 0; i < len(actual); i++ {
			areTheSame = areTheSame && actual[i] == expected[i]
		}
	}

	if !areTheSame {
		t.Errorf("Input=%v k=%d Expected=%v Actual=%v", words, k, expected, actual)
	}
}
