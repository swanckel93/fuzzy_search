package search

import (
	"strings"
	"sync"

	"github.com/agnivade/levenshtein"
)

// SearchResult represents a fuzzy match result
type SearchResult struct {
	Sentence string `json:"sentence"`
	Index    int    `json:"index"`
	Match    string `json:"match"`
	Distance int    `json:"distance"`
}

// FuzzySearch performs a fuzzy search for a query in a slice of sentences using goroutines.
func FuzzySearch(query string, sentences []string) []SearchResult {
	var wg sync.WaitGroup
	resultsChan := make(chan SearchResult, len(sentences))

	for i, sentence := range sentences {
		wg.Add(1)

		go func(idx int, s string) {
			defer wg.Done()
			bestMatch, bestIndex, bestDist := findBestFuzzyMatch(query, s)
			if bestMatch != "" {
				resultsChan <- SearchResult{
					Sentence: s,
					Index:    bestIndex,
					Match:    bestMatch,
					Distance: bestDist,
				}
			}
		}(i, sentence)
	}

	wg.Wait()
	close(resultsChan)

	var results []SearchResult
	for result := range resultsChan {
		results = append(results, result)
	}

	return results
}

// findBestFuzzyMatch finds the substring in `sentence` that best matches `query` using Levenshtein distance.
func findBestFuzzyMatch(query, sentence string) (match string, index int, distance int) {
	queryLen := len(query)
	sentenceLower := strings.ToLower(sentence)
	queryLower := strings.ToLower(query)

	bestDistance := -1
	bestMatch := ""
	bestIndex := -1

	for i := 0; i <= len(sentenceLower)-queryLen; i++ {
		substr := sentenceLower[i : i+queryLen]
		dist := levenshtein.ComputeDistance(substr, queryLower)
		if bestDistance == -1 || dist < bestDistance {
			bestDistance = dist
			bestMatch = sentence[i : i+queryLen]
			bestIndex = i
		}
	}

	return bestMatch, bestIndex, bestDistance
}
