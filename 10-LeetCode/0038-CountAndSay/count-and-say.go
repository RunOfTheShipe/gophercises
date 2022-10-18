package countandsay

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else {
		var strToSay = countAndSay(n - 1)
		return strToSay
	}
}
