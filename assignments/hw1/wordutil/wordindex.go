package wordutil

import (
	"strings"
)

// Finds first occurrence of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// the index of the first occurence of the word in the input string as the
// corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndex(s string) map[string]int {
	// TODO: implement me
	words := strings.Fields(s)
	indexMap := make(map[string]int)
	const BLANK int = -50
	for _, word := range words {
		word = strings.ToLower(word)
		indexMap[word] = BLANK
	}
	for _, word := range words {
		word = strings.ToLower(word)
		for indexMainStr := range s {
			isBeginningOfTheWord := MatchesInWordAtIndex(s, word, indexMainStr, false)
			if isBeginningOfTheWord && indexMap[word] == BLANK {
				indexMap[word] = indexMainStr
			}
		}

	}
	return indexMap
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
}
func MatchesInWordAtIndex(str, substring string, index int, caseSensitive bool) bool {
	upUntilIndex := str[index:]
	if !caseSensitive {
		upUntilIndex = strings.ToLower(upUntilIndex)
	}
	return strings.HasPrefix(upUntilIndex, substring)
}
