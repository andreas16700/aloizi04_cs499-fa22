package wordutil

import (
	"strings"
)

// Finds all occurrences of each word in a string.
//
// Returns a map that stores each unique word in the string s as the key and
// a slice that contains the index of each occurence of the word in the input
// string as the corresponding value.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the
// same word.
func WordIndexAll(s string) map[string][]int {
	words := uniqueWords(s, true)
	indexMap := make(map[string][]int)
	for _, word := range words {
		word = strings.ToLower(word)
		for indexMainStr := range s {
			isBeginningOfTheWord := MatchesInWordAtIndex(s, word, indexMainStr, false)
			if isBeginningOfTheWord {
				arrayOfBeginningIndices := indexMap[word]
				if arrayOfBeginningIndices == nil {
					arrayOfBeginningIndices = make([]int, 0)
				}
				arrayOfBeginningIndices = append(arrayOfBeginningIndices, indexMainStr)
				indexMap[word] = arrayOfBeginningIndices
			}
		}
	}
	return indexMap
}
func uniqueWords(s string, convertToLowercase bool) []string {
	words := strings.Fields(s)
	wordMap := make(map[string]bool)
	for _, word := range words {
		if convertToLowercase {
			word = strings.ToLower(word)
		}
		wordMap[word] = true
	}
	keys := make([]string, len(wordMap))
	i := 0
	for k := range wordMap {
		keys[i] = k
		i++
	}
	return keys
}
