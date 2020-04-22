// Package stringcase implements utility functions for
// changing the case of simple strings.
package stringcase

import (
	"strings"
)

// ToSnake converts a string to snake_case.
func ToSnake(s string) string {
	return joinWords(splitWords(s), "_", lowerWordTransformer)
}

// ToUpperSnake converts a string to UPPER_SNAKE_CASE.
func ToUpperSnake(s string) string {
	return joinWords(splitWords(s), "_", upperWordTransformer)
}

// ToCamel converts a string to camelCase.
func ToCamel(s string) string {
	return joinWords(splitWords(s), "", func(i int, word string) string {
		if len(word) == 0 {
			return ""
		}
		if i == 0 {
			return strings.ToLower(word)
		}
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	})
}

// ToUpperCamel converts a string to UpperCamelCase (aka PascalCase).
func ToUpperCamel(s string) string {
	return joinWords(splitWords(s), "", func(_ int, word string) string {
		if len(word) == 0 {
			return ""
		}
		return strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
	})
}

// ToKebab converts a string to kebab-case.
func ToKebab(s string) string {
	return joinWords(splitWords(s), "-", lowerWordTransformer)
}

// ToUpperKebab converts a string to UPPER-KEBAB-CASE.
func ToUpperKebab(s string) string {
	return joinWords(splitWords(s), "-", upperWordTransformer)
}
