package slugify

import (
	"regexp"
	"slices"
	"strings"
	"unicode"
)

// Options - configures the behavior of slug generation.
type Options struct {
	// character(s) used to replace spaces and separators (default: "-")
	Separator string

	// maxmium length of the resulting slug, 0 means no limit
	MaxLength int

	// map of custom character replacements to apply
	CustomReplacements map[string]string
}

var alphabet = map[rune]string{
	'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d",
	'е': "e", 'ё': "yo", 'ж': "zh", 'з': "z", 'и': "i",
	'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
	'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t",
	'у': "u", 'ф': "f", 'х': "h", 'ц': "ts", 'ч': "ch",
	'ш': "sh", 'щ': "shch", 'ы': "y",
	'э': "e", 'ю': "yu", 'я': "ya"}

// Make - converts a string to a URL-friendly slug using default options.
// It transliterates Cyrillic characters, replaces separators with hyphens,
// removes file extensions, and converts to lowercase.
func Make(source string) string {
	return MakeWithOptions(source, Options{})
}

// MakeWithOptions - converts a string to a URL-friendly slug with custom options.
// It processes the input string by transliterating Cyrillic characters to Latin,
// replacing separators and special characters, and applying length limits.
//
// The function preserves Latin letters and digits, transliterates Cyrillic characters
// using the internal alphabet mapping, applies custom replacements if specified,
// and ensures no consecutive or trailing separators in the output.
func MakeWithOptions(source string, options Options) string {
	var stringAccum strings.Builder
	var prevChar string

	formattedSource := formatString(source)
	regExp, _ := regexp.Compile("[a-z]")
	maxLength := options.MaxLength
	separator := "-"

	if options.Separator != "" {
		separator = options.Separator
	}

	for _, el := range formattedSource {
		transformedLetter, isInAlphabet := alphabet[el]
		replacement, isReplacement := options.CustomReplacements[string(el)]
		isLatin := regExp.MatchString(string(el))
		isDigit := unicode.IsDigit(el)

		if isSeparator(el) || isReplacement {
			if prevChar != separator && prevChar != "" {
				stringAccum.WriteString(separator)
			}
			prevChar = separator
		}

		if isInAlphabet {
			stringAccum.WriteString(transformedLetter)
			prevChar = string(transformedLetter[len(transformedLetter)-1])
		}

		if isDigit || isLatin {
			stringAccum.WriteRune(el)
			prevChar = string(el)
		}

		if isReplacement {
			stringAccum.WriteString(replacement)
			prevChar = string(replacement[len(replacement)-1])
		}

	}

	rawResult := stringAccum.String()

	if maxLength > 0 && maxLength < len(rawResult) {
		rawResult = rawResult[:maxLength]
	}

	if len(rawResult) > 0 && separator != "" && strings.HasSuffix(rawResult, separator) {
		rawResult = rawResult[:len(rawResult)-1]
	}

	return rawResult
}

func isSeparator(char rune) bool {
	separators := []rune{' ', '.', '-', '/', '_'}

	return slices.Contains(separators, char)
}

func removeExtension(fileName string) string {
	extension := strings.LastIndex(fileName, ".")

	if extension > 0 {
		return fileName[:extension]
	}

	return fileName

}

func formatString(inputString string) string {
	lowered := strings.ToLower(inputString)
	trimmed := strings.TrimSpace(lowered)
	result := removeExtension(trimmed)

	return result
}
