package wordutil

import (
	"strings"
)

// Counts occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// its count as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordCount(s string) map[string]int {
	wordsSlice := strings.Fields(s)
	wordMap := make(map[string]int)
	for _, word := range wordsSlice {
		word = strings.ToLower(word)
		wordMap[word]++
	}
	return wordMap
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
}
