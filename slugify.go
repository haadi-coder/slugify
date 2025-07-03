package slugify

import (
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Options - configures the behavior of slug generation.
type Options struct {
	// Seperator character(s) used to replace spaces and separators (default: "-")
	Separator string

	// Maxlength maxmium length of the resulting slug, 0 means no limit
	MaxLength int

	// CustomReplacements map of custom character replacements to apply
	CustomReplacements map[rune]string
}

var specChars = []rune{' ', '.', '-', '/', '_'}

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
func Make(s string) string {
	return MakeWithOptions(s, Options{})
}

// MakeWithOptions - converts a string to a URL-friendly slug with custom options.
// It processes the input string by transliterating Cyrillic characters to Latin,
// replacing separators and special characters, and applying length limits.
//
// The function preserves Latin letters and digits, transliterates Cyrillic characters
// using the internal alphabet mapping, applies custom replacements if specified,
// and ensures no consecutive or trailing separators in the output.
func MakeWithOptions(s string, opts Options) string {
	var sb strings.Builder
	var prevChar string
	var runeCount int

	prepared := prepareString(s)
	separator := determineSeparator(opts.Separator)

	for _, r := range prepared {
		if opts.MaxLength > 0 && runeCount >= opts.MaxLength {
			break
		}

		if isSeparator(r) {
			if prevChar != separator {
				sb.WriteString(separator)
				runeCount++
			}
			prevChar = separator
			continue
		}

		if replacement, ok := opts.CustomReplacements[r]; ok {

			if prevChar != separator {
				sb.WriteString(separator)
				runeCount++
			}

			sb.WriteString(replacement)
			prevChar = string(replacement[len(replacement)-1])
			runeCount += utf8.RuneCountInString(replacement)
			continue
		}

		if replacement, ok := alphabet[r]; ok {
			sb.WriteString(replacement)
			prevChar = string(replacement[len(replacement)-1])
			runeCount += utf8.RuneCountInString(replacement)
			continue
		}

		if unicode.IsDigit(r) || unicode.Is(unicode.Latin, r) {
			sb.WriteRune(r)
			prevChar = string(r)
			runeCount++
			continue
		}
	}

	return sb.String()
}

func isSeparator(char rune) bool {
	return slices.Contains(specChars, char)
}

func determineSeparator(separator string) string {

	if separator != "" {
		return separator
	}

	return "-"
}

func prepareString(inputString string) string {
	lowered := strings.ToLower(inputString)
	trimmed := strings.Trim(lowered, string(specChars))

	return trimmed
}
