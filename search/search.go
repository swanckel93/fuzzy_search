package search

import (
	"fmt"
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
	"github.com/swanckel93/fuzzy_api/models"
)

func FuzzySearch(query string, sentences []string) []models.SearchResult {
	type resultWithScore struct {
		models.SearchResult
	}

	var results []resultWithScore
	loweredQuery := strings.ToLower(query)

	fmt.Printf("Starting fuzzy search for query: %s\n", loweredQuery)

	for i, sentence := range sentences {
		loweredSentence := strings.ToLower(sentence)

		// Check if the sentence contains the query word
		if strings.Contains(loweredSentence, loweredQuery) {
			// If it contains the word, set Levenshtein distance to 0
			fmt.Printf("Sentence %d: \"%s\" contains the query. Distance set to 0.\n", i, sentence)
			results = append(results, resultWithScore{
				SearchResult: models.SearchResult{
					Sentence: sentence,
					Index:    i,
					Match:    query,
					Distance: 0,
				},
			})
		} else {
			// If it doesn't contain the word, calculate the Levenshtein distance
			distance := levenshtein.ComputeDistance(loweredQuery, loweredSentence)
			fmt.Printf("Sentence %d: \"%s\"\n", i, sentence)
			fmt.Printf("Lowered: \"%s\"\n", loweredSentence)
			fmt.Printf("Levenshtein distance: %d\n", distance)

			if distance < len(loweredQuery) {
				fmt.Println("-> Accepted (distance threshold passed)")
				results = append(results, resultWithScore{
					SearchResult: models.SearchResult{
						Sentence: sentence,
						Index:    i,
						Match:    query,
						Distance: distance,
					},
				})
			} else {
				fmt.Println("-> Skipped (distance too high)")
			}
		}
	}

	// Sort results by distance
	sort.Slice(results, func(i, j int) bool {
		return results[i].Distance < results[j].Distance
	})

	// Take top 10 results
	top := []models.SearchResult{}
	for i := 0; i < len(results) && i < 10; i++ {
		top = append(top, results[i].SearchResult)
	}

	fmt.Printf("Returning %d result(s)\n", len(top))
	return top
}
