package search

import (
	"testing"
)

func TestFuzzySearch(t *testing.T) {
	tests := []struct {
		name      string
		query     string
		sentences []string
		expected  func([]SearchResult) bool
	}{
		{
			name:      "Exact match",
			query:     "hello",
			sentences: []string{"hello world", "say hello"},
			expected: func(results []SearchResult) bool {
				for _, res := range results {
					if res.Distance == 0 && res.Match == "hello" {
						return true
					}
				}
				return false
			},
		},
		{
			name:      "Single character off",
			query:     "hella",
			sentences: []string{"hello world", "no match here"},
			expected: func(results []SearchResult) bool {
				for _, res := range results {
					if res.Match == "hello" && res.Distance == 1 {
						return true
					}
				}
				return false
			},
		},
		{
			name:      "No close match",
			query:     "xyz",
			sentences: []string{"this is a test", "another sentence"},
			expected: func(results []SearchResult) bool {
				return len(results) == 2 // Still returns best matches, even if not great
			},
		},
		{
			name:      "Mixed results",
			query:     "cat",
			sentences: []string{"the cat sat", "a caterpillar", "bat and rat"},
			expected: func(results []SearchResult) bool {
				return len(results) == 3
			},
		},
		{
			name:      "Empty sentence list",
			query:     "test",
			sentences: []string{},
			expected: func(results []SearchResult) bool {
				return len(results) == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FuzzySearch(tt.query, tt.sentences)
			if !tt.expected(got) {
				t.Errorf("Test %q failed. Got: %+v", tt.name, got)
			}
		})
	}
}

func TestFuzzySearchOrdering(t *testing.T) {
	query := "cat"
	sentences := []string{
		"The cat is on the mat",
		"A catalog of items",
		"That was a catastrophe",
		"Cats are cute",
		"Concatenate the strings",
	}

	results := FuzzySearch(query, sentences)

	if len(results) < 2 {
		t.Fatalf("Expected at least 2 results, got %d", len(results))
	}

	for i := 0; i < len(results)-1; i++ {
		a := results[i]
		b := results[i+1]

		if a.Distance > b.Distance {
			t.Errorf("Results not sorted by distance at index %d: got %d > %d", i, a.Distance, b.Distance)
		} else if a.Distance == b.Distance && a.Index > b.Index {
			t.Errorf("Results not sorted by index when distances equal at index %d: got %d > %d", i, a.Index, b.Index)
		}
	}
}