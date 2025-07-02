package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

type Options struct {
	Separator          string
	MaxLength          int
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

func main() {
	output := MakeWithOptions("C++", Options{CustomReplacements: map[string]string{"+": "plus"}})
	fmt.Println(output)
}

func Make(source string) string {
	return MakeWithOptions(source, Options{})
}

func MakeWithOptions(inputString string, options Options) string {
	var stringAccum strings.Builder
	var prevChar string

	formattedInputString := formatString(inputString)
	regExp, _ := regexp.Compile("[a-z]")
	maxLength := options.MaxLength
	separator := "-"

	if options.Separator != "" {
		separator = options.Separator
	}

	for _, el := range formattedInputString {
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
