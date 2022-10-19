package frequentwords

import "sort"

// Could you solve it in O(n log(k)) time and O(n) extra space?

func topKFrequent(words []string, k int) []string {

	// build a map of words and counts (O(n))
	var wordCounts map[string]int = make(map[string]int) // O(n)
	for _, word := range words {
		var _, found = wordCounts[word]
		if !found {
			wordCounts[word] = 0
		}
		wordCounts[word]++
	}

	// build a second map of counts to set of words (O(n) - worst case being all unique words)
	var countWords map[int][]string = make(map[int][]string) // O(n)
	for word, count := range wordCounts {
		var _, found = countWords[count]
		if !found {
			countWords[count] = []string{}
		}

		countWords[count] = append(countWords[count], word)
	}

	// in a "worst case," the input string could contain the same
	// word repeated for the entire duration of the string
	var frequent []string = []string{}
	for i := len(words); i > 0; i-- {
		var wordsWithCount, found = countWords[i]
		if found {
			// sort these words (O(n log(n)) - all words in the list could be unique)
			sort.Strings(wordsWithCount)

			// found words with a specific count, take the resulting
			// words from this list and append to the final list

			for _, word := range wordsWithCount {
				if k > 0 {
					frequent = append(frequent, word)
					k--
				}
			}
		}
	}

	return frequent
}
