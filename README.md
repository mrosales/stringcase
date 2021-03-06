# stringcase

[![Build][build-badge]][build]
[![codecov][codecov-badge]][codecov]
[![goreport][goreport-badge]][goreport]
[![GoDoc][godoc-badge]][godoc]
[![MIT License][license-badge]][license]

Easy string case conversion in Go

## Usage

```go
s := "MIXED-caseString.test"
```

| Converter           | Output                     |
|---------------------|----------------------------|
| `ToSnake(s)`        | `mixed_case_string_test`   |
| `ToUpperSnake(s)`   | `MIXED_CASE_STRING_TEST`   |
| `ToCamel(s)`        | `mixedCaseStringTest`      |
| `ToUpperCamel(s)`   | `MixedCaseStringTest`      |
| `ToKebab(s)`        | `mixed-case-string-test`   |
| `ToUpperKebab(s)`   | `MIXED-CASE-STRING-TEST`   |

If you are curious about the rules for converting strings, take a look at
[the test file](stringcase_test.go) for examples.

## Install

This package uses go modules and you should be able to use the latest release tag.

```shell
go get -u github.com/mrosales/stringcase
```

## Implementation Details

### Special Characters

All converters will **remove any non-alphanumeric characters**. That means
things like emojis, punctuation, whitespace and other unicode characters
will all be removed.

### Word Breaks

Before transforming the string, it is first split into words. This
implementation tries to do the **right thing**™, but it may very well
break some part of your use case. If you encounter something that seems
**incorrect**, don't hesitate to file an issue. That being said, it may
not be possible to update the library to cover every edge case.

The following cases trigger a word break:

* Any non alphanumeric character (`[^a-zA-Z0-9]`)
* Any whitespace (`\s+`)
* Transitioning from lowercase to uppercase
  This attempts to respect Go's preference to capitalize acronyms and abbreviations.

  * `fooBar` -> `foo Bar`
  * `fooBAR` -> `foo BAR`
  * `userID` -> `user ID`
  * `userIDFOO` -> `user IDFOO` (*nothing we can really do about a string like this*)

* Transitioning from uppercase to lowercase
  * `userIDBar` -> `user ID Bar`

## Inspiration

* [`github.com/iancoleman/strcase`](https://github.com/iancoleman/strcase)
  * This is a great library that I used for some time before this, but I felt
    like the camelcase implementation in particular seemed inconsistent.
    For example, I would prefer `ToCamel("fooBAR")` to return `"fooBar"` rather than
    `"fooBar"`. This is not necessarily wrong, but it does not match how
    `ToSnake("fooBAR")` for that library returns `foo_bar`.
  * The change that this `stringcase` library makes is to first split the string
    into words and then join them in canonical fashion for each type of casing.
    This is slightly less efficient, but I find it to produce more predictable
    results.

---

Released under the [MIT License](LICENSE.txt).

[build]: https://github.com/mrosales/stringcase/actions?query=branch%3Amaster
[build-badge]: https://img.shields.io/github/workflow/status/mrosales/stringcase/Test/master?style=flat-square
[codecov]: https://codecov.io/gh/mrosales/stringcase
[codecov-badge]: https://img.shields.io/codecov/c/gh/mrosales/stringcase?style=flat-square
[goreport]: https://goreportcard.com/report/github.com/mrosales/stringcase
[goreport-badge]: https://goreportcard.com/badge/github.com/mrosales/stringcase?style=flat-square
[godoc]: https://pkg.go.dev/github.com/mrosales/stringcase
[godoc-badge]: https://img.shields.io/badge/godoc-reference-blue?style=flat-square
[license]: LICENSE.txt
[license-badge]: https://img.shields.io/github/license/mrosales/stringcase?style=flat-square
