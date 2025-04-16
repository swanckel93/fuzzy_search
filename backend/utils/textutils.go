package utils

import (
	"regexp"
	"strings"
)

var sentenceRegex = regexp.MustCompile(`(?m)([^.!?]*[.!?])`)

func SplitIntoSentences(text string) []string {
	matches := sentenceRegex.FindAllString(text, -1)
	sentences := make([]string, 0, len(matches))
	for _, s := range matches {
		trimmed := strings.TrimSpace(s)
		if trimmed != "" {
			sentences = append(sentences, trimmed)
		}
	}
	return sentences
}

func HighlightMatch(sentence, match string) string {
	if match == "" {
		return sentence
	}
	return strings.ReplaceAll(sentence, match, "<mark>"+match+"</mark>")
}
