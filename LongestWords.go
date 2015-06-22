package codeTest

import (
	"strings"
)

func LongestWords(words string) []string {
	results := []string{}
	maxWordSize := 0
	for _, word := range strings.Split(words, " ") {
		if len(word) > maxWordSize {
			results = []string{word}
			maxWordSize = len(word)
		} else if len(word) == maxWordSize {
			found := false
			for _, search := range results {
				if strings.ToLower(search) == strings.ToLower(word) {
					found = true
					break
				}
			}
			if !found {
				results = append(results, word)
			}
		}
	}
	return results
}
