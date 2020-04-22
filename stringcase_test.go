package stringcase

import (
	"fmt"
	"testing"
)

func ExampleToSnake() {
	fmt.Println(ToSnake("snake case"))
	// Output: snake_case
}

func TestToSnake(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"kebab-case", "kebab_case"},
		{"snake_case", "snake_case"},
		{"UPPER_SNAKE_CASE", "upper_snake_case"},
		{"UPPER-KEBAB-CASE", "upper_kebab_case"},
		// more interesting cases
		{"email@foo.com", "email_foo_com"},
		{"no emojis allowed 👎", "no_emojis_allowed"},
		{"other/characters(should*&^%be!removed", "other_characters_should_be_removed"},
		{"numbers_dont_split1", "numbers_dont_split1"},
		{"dot.delimited.words", "dot_delimited_words"},
		{"fooID", "foo_id"},
		{"fooIDBar", "foo_id_bar"},
		{"JSONStuff", "json_stuff"},
		{" leading whitespace", "leading_whitespace"},
		{"trailing whitespace ", "trailing_whitespace"},
		{"Many Capitalized Words", "many_capitalized_words"},
		{"Very-veryConfusedANDOdd_CASED string", "very_very_confused_and_odd_cased_string"},
		// edge case regression tests
		{"fooBBAR", "foo_bbar"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToSnake(tt.input); got != tt.want {
				t.Errorf("ToSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToUpperSnake() {
	fmt.Println(ToUpperSnake("upper snake case"))
	// Output: UPPER_SNAKE_CASE
}

func TestToUpperSnake(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "CAMEL_CASE"},
		{"PascalCase", "PASCAL_CASE"},
		{"kebab-case", "KEBAB_CASE"},
		{"snake_case", "SNAKE_CASE"},
		{"UPPER_SNAKE_CASE", "UPPER_SNAKE_CASE"},
		{"UPPER-KEBAB-CASE", "UPPER_KEBAB_CASE"},
		// more interesting cases
		{"email@foo.com", "EMAIL_FOO_COM"},
		{"no emojis allowed 👎", "NO_EMOJIS_ALLOWED"},
		{"other/characters(should*&^%be!removed", "OTHER_CHARACTERS_SHOULD_BE_REMOVED"},
		{"numbers_dont_split1", "NUMBERS_DONT_SPLIT1"},
		{"dot.delimited.words", "DOT_DELIMITED_WORDS"},
		{"fooID", "FOO_ID"},
		{"fooIDBar", "FOO_ID_BAR"},
		{"JSONStuff", "JSON_STUFF"},
		{" leading whitespace", "LEADING_WHITESPACE"},
		{"trailing whitespace ", "TRAILING_WHITESPACE"},
		{"Many Capitalized Words", "MANY_CAPITALIZED_WORDS"},
		{"Very-veryConfusedANDOdd_CASED string", "VERY_VERY_CONFUSED_AND_ODD_CASED_STRING"},
		// edge case regression tests
		{"fooBBAR", "FOO_BBAR"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToUpperSnake(tt.input); got != tt.want {
				t.Errorf("ToUpperSnake() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToCamel() {
	fmt.Println(ToCamel("camel case"))
	// Output: camelCase
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "camelCase"},
		{"PascalCase", "pascalCase"},
		{"kebab-case", "kebabCase"},
		{"snake_case", "snakeCase"},
		{"UPPER_SNAKE_CASE", "upperSnakeCase"},
		{"UPPER-KEBAB-CASE", "upperKebabCase"},
		// more interesting cases
		{"email@foo.com", "emailFooCom"},
		{"no emojis allowed 👎", "noEmojisAllowed"},
		{"other/characters(should*&^%be!removed", "otherCharactersShouldBeRemoved"},
		{"numbers_dont_split1", "numbersDontSplit1"},
		{"dot.delimited.words", "dotDelimitedWords"},
		{"fooID", "fooId"},
		{"fooIDBar", "fooIdBar"},
		{"JSONStuff", "jsonStuff"},
		{" leading whitespace", "leadingWhitespace"},
		{"trailing whitespace ", "trailingWhitespace"},
		{"Many Capitalized Words", "manyCapitalizedWords"},
		{"Very-veryConfusedANDOdd_CASED string", "veryVeryConfusedAndOddCasedString"},
		// edge case regression tests
		{"fooBBAR", "fooBbar"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToCamel(tt.input); got != tt.want {
				t.Errorf("ToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToUpperCamel() {
	fmt.Println(ToUpperCamel("upper_camel_case"))
	// Output:
	// UpperCamelCase
}

func TestToUpperCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "CamelCase"},
		{"PascalCase", "PascalCase"},
		{"kebab-case", "KebabCase"},
		{"snake_case", "SnakeCase"},
		{"UPPER_SNAKE_CASE", "UpperSnakeCase"},
		{"UPPER-KEBAB-CASE", "UpperKebabCase"},
		// more interesting cases
		{"email@foo.com", "EmailFooCom"},
		{"no emojis allowed 👎", "NoEmojisAllowed"},
		{"other/characters(should*&^%be!removed", "OtherCharactersShouldBeRemoved"},
		{"numbers_dont_split1", "NumbersDontSplit1"},
		{"dot.delimited.words", "DotDelimitedWords"},
		{"fooID", "FooId"},
		{"fooIDBar", "FooIdBar"},
		{"JSONStuff", "JsonStuff"},
		{" leading whitespace", "LeadingWhitespace"},
		{"trailing whitespace ", "TrailingWhitespace"},
		{"Many Capitalized Words", "ManyCapitalizedWords"},
		{"Very-veryConfusedANDOdd_CASED string", "VeryVeryConfusedAndOddCasedString"},
		// edge case regression tests
		{"fooBBAR", "FooBbar"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToUpperCamel(tt.input); got != tt.want {
				t.Errorf("ToUpperCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToKebab() {
	fmt.Println(ToKebab("kebab case"))
	// Output: kebab-case
}

func TestToKebab(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "camel-case"},
		{"PascalCase", "pascal-case"},
		{"kebab-case", "kebab-case"},
		{"snake_case", "snake-case"},
		{"UPPER_SNAKE_CASE", "upper-snake-case"},
		{"UPPER-KEBAB-CASE", "upper-kebab-case"},
		// more interesting cases
		{"email@foo.com", "email-foo-com"},
		{"no emojis allowed 👎", "no-emojis-allowed"},
		{"other/characters(should*&^%be!removed", "other-characters-should-be-removed"},
		{"numbers_dont_split1", "numbers-dont-split1"},
		{"dot.delimited.words", "dot-delimited-words"},
		{"fooID", "foo-id"},
		{"fooIDBar", "foo-id-bar"},
		{"JSONStuff", "json-stuff"},
		{" leading whitespace", "leading-whitespace"},
		{"trailing whitespace ", "trailing-whitespace"},
		{"Many Capitalized Words", "many-capitalized-words"},
		{"Very-veryConfusedANDOdd_CASED string", "very-very-confused-and-odd-cased-string"},
		// edge case regression tests
		{"fooBBAR", "foo-bbar"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToKebab(tt.input); got != tt.want {
				t.Errorf("ToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleToUpperKebab() {
	fmt.Println(ToUpperKebab("upper kebab case"))
	// Output: UPPER-KEBAB-CASE
}

func TestToUpperKebab(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		// standard tests
		{"", ""},
		{"camelCase", "CAMEL-CASE"},
		{"PascalCase", "PASCAL-CASE"},
		{"kebab-case", "KEBAB-CASE"},
		{"snake_case", "SNAKE-CASE"},
		{"UPPER_SNAKE_CASE", "UPPER-SNAKE-CASE"},
		{"UPPER-KEBAB-CASE", "UPPER-KEBAB-CASE"},
		// more interesting cases
		{"email@foo.com", "EMAIL-FOO-COM"},
		{"no emojis allowed 👎", "NO-EMOJIS-ALLOWED"},
		{"other/characters(should*&^%be!removed", "OTHER-CHARACTERS-SHOULD-BE-REMOVED"},
		{"numbers_dont_split1", "NUMBERS-DONT-SPLIT1"},
		{"dot.delimited.words", "DOT-DELIMITED-WORDS"},
		{"fooID", "FOO-ID"},
		{"fooIDBar", "FOO-ID-BAR"},
		{"JSONStuff", "JSON-STUFF"},
		{" leading whitespace", "LEADING-WHITESPACE"},
		{"trailing whitespace ", "TRAILING-WHITESPACE"},
		{"Many Capitalized Words", "MANY-CAPITALIZED-WORDS"},
		{"Very-veryConfusedANDOdd_CASED string", "VERY-VERY-CONFUSED-AND-ODD-CASED-STRING"},
		// edge case regression tests
		{"fooBBAR", "FOO-BBAR"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := ToUpperKebab(tt.input); got != tt.want {
				t.Errorf("ToUpperKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}
