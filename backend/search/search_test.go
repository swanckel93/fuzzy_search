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
