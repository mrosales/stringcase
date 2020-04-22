package stringcase

import (
	"regexp"
	"strings"
)

var (
	// match any non alphanumeric characters
	invalidCharacterPattern = regexp.MustCompile(`[^a-zA-Z0-9]`)

	// detects when the case changes from upper to lower
	multiWordPattern = regexp.MustCompile(`([A-Z0-9]+)([A-Z][a-z0-9]+)`)

	// match when case changes from lower to upper
	capChangePattern = regexp.MustCompile(`([a-z0-9])([A-Z])`)

	// clean up multiple consecutive spaces
	duplicateSpacePattern = regexp.MustCompile(`\s+`)
)

func splitWords(s string) []string {
	s = invalidCharacterPattern.ReplaceAllLiteralString(s, " ")
	s = multiWordPattern.ReplaceAllString(s, "$1 $2")
	s = capChangePattern.ReplaceAllString(s, "$1 $2")
	s = duplicateSpacePattern.ReplaceAllLiteralString(s, " ")
	s = strings.TrimSpace(s)
	return strings.Split(s, " ")
}

func joinWords(words []string, delimiter string, wordTransformer func(int, string) string) string {
	for i, w := range words {
		words[i] = wordTransformer(i, w)
	}
	return strings.Join(words, delimiter)
}

func lowerWordTransformer(_ int, word string) string {
	return strings.ToLower(word)
}

func upperWordTransformer(_ int, word string) string {
	return strings.ToUpper(word)
}
