package countandsay

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		var strToSay = countAndSay(n - 1)

		var saidString string = ""

		var currentRune rune = rune(strToSay[0])
		var runeCount int = 1

		for _, r := range strToSay[1:] {
			if currentRune == r {
				// same rune - just increment the count
				runeCount++
			} else {
				// new rune found, append the old count+rune to
				// the said string and start counting again
				saidString = fmt.Sprintf("%s%d%s", saidString, runeCount, string(currentRune))

				// restart counting
				currentRune = r
				runeCount = 1
			}
		}

		// add the last rune/count to the said string
		saidString = fmt.Sprintf("%s%d%s", saidString, runeCount, string(currentRune))

		return saidString
	}
}
