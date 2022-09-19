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
	// TODO: implement me
	// HINT: You may find the `strings.Index` and `strings.ToLower` functions helpful
	// TODO: implement me
	words := strings.Fields(s)
	indexMap := make(map[string][]int)
	for _, word := range words{
		word = strings.ToLower(word)
		for indexMainStr, _ := range s{
			isBeginningOfTheWord := matchesInWordAtIndex(s, word, indexMainStr)
			if isBeginningOfTheWord{
				arrayOfBeginningIndices := indexMap[word]
				if arrayOfBeginningIndices == nil{
					arrayOfBeginningIndices = make([]int, 5)
				}
				arrayOfBeginningIndices = append(arrayOfBeginningIndices, indexMainStr)
				indexMap[word]=arrayOfBeginningIndices
			}
		}
	}
	return indexMap
}
