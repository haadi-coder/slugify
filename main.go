package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode"
)

var alphabet = map[rune]string{'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d",
	'е': "e", 'ё': "yo", 'ж': "zh", 'з': "z", 'и': "i",
	'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
	'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t",
	'у': "u", 'ф': "f", 'х': "h", 'ц': "ts", 'ч': "ch",
	'ш': "sh", 'щ': "shch", 'ы': "y",
	'э': "e", 'ю': "yu", 'я': "ya"}

func main() {
	output := Make("---Hello World---")
	fmt.Println(output)
}

func Make(source string) string {
	var stringAccum strings.Builder
	var result string
	var prevChar rune

	formattedSource := removeExtension(strings.TrimSpace(strings.ToLower(source)))
	regExp, _ := regexp.Compile("[a-z]")

	for _, el := range formattedSource {
		transformedLetter, isInAlphabet := alphabet[el]
		isLatin := regExp.MatchString(string(el))
		isDigit := unicode.IsDigit(el)

		if isSeparator(el) {
			if prevChar != '-' && prevChar != 0 {
				stringAccum.WriteRune('-')
			}
			prevChar = '-'
		}

		if isInAlphabet {
			stringAccum.WriteString(transformedLetter)
			prevChar = rune(transformedLetter[len(transformedLetter)-1])
		}

		if isDigit || isLatin {
			stringAccum.WriteRune(el)
			prevChar = el
		}

	}

	result = stringAccum.String()

	if len(result) > 0 && result[len(result)-1] == '-' {
		result = result[:len(result)-1]
	}

	return result
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
