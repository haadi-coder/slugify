package slugify

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Options - configures the behavior of slug generation.
type Options struct {
	// character(s) used to replace spaces and separators (default: "-")
	Separator rune

	// maxmium length of the resulting slug, 0 means no limit
	MaxLength int

	// map of custom character replacements to apply
	CustomReplacements map[rune]string
}

var separators = []rune{' ', '.', '-', '/', '_'}

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
	var prevChar rune
	var runeCount int

	preparedSource := prepareString(source)
	separator := determineSeparator(options.Separator)

	for _, el := range preparedSource {

		if options.MaxLength > 0 && runeCount >= options.MaxLength {
			break
		}

		if isSeparator(el) {
			if prevChar != separator {
				stringAccum.WriteRune(separator)
				runeCount++
			}
			prevChar = separator
			continue
		}

		if replacement, ok := options.CustomReplacements[el]; ok {

			if prevChar != separator {
				stringAccum.WriteRune(separator)
				runeCount++
			}

			stringAccum.WriteString(replacement)
			prevChar = rune(replacement[len(replacement)-1])
			runeCount += utf8.RuneCountInString(replacement)
			continue
		}

		if transformedLetter, isInAlphabet := alphabet[el]; isInAlphabet {
			stringAccum.WriteString(transformedLetter)
			prevChar = rune(transformedLetter[len(transformedLetter)-1])
			runeCount += utf8.RuneCountInString(transformedLetter)
			continue
		}

		if unicode.IsDigit(el) || unicode.Is(unicode.Latin, el) {
			stringAccum.WriteRune(el)
			prevChar = el
			runeCount++
			continue
		}

	}

	return stringAccum.String()

}

func isSeparator(char rune) bool {
	return slices.Contains(separators, char)
}

func determineSeparator(separator rune) rune {

	if separator != 0 {
		return separator
	}

	return '-'
}

func removeSuffix(fileName string) string {
	suffix := strings.LastIndex(fileName, ".")

	if suffix > 0 {
		return fileName[:suffix]
	}

	return fileName
}

func prepareString(inputString string) string {
	lowered := strings.ToLower(inputString)
	trimmed := strings.Trim(lowered, string(separators))
	result := removeSuffix(trimmed)

	return result
}
