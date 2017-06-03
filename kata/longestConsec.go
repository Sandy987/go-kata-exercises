// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go - An attempt to solve this kata with Go

package kata

// LongestConsec returns the longest possible concatenation of k consecuctive words in strarr
func LongestConsec(strarr []string, k int) string {
	currentBiggestWord := ""

	if len(strarr) == 0 || k <= 0 || k > len(strarr) { // Base cases for when there are no elements to check
		return currentBiggestWord
	}

	for i := 0; i < len(strarr)-k+1; i++ {
		word := ""
		for j := 0; j < k; j++ {
			word += strarr[i+j]
		}
		if len(word) > len(currentBiggestWord) {
			currentBiggestWord = word
		}
	}
	return currentBiggestWord
}
