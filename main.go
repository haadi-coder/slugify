package main

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var alphabet = map[rune]string{'а': "a", 'б': "b", 'в': "v", 'г': "g", 'д': "d",
	'е': "e", 'ё': "yo", 'ж': "zh", 'з': "z", 'и': "i",
	'й': "y", 'к': "k", 'л': "l", 'м': "m", 'н': "n",
	'о': "o", 'п': "p", 'р': "r", 'с': "s", 'т': "t",
	'у': "u", 'ф': "f", 'х': "h", 'ц': "ts", 'ч': "ch",
	'ш': "sh", 'щ': "shch", 'ъ': "", 'ы': "y", 'ь': "",
	'э': "e", 'ю': "yu", 'я': "ya"}

func main() {
	output := Make("")
	fmt.Println(output)
}

func Make(source string) string {
	var result strings.Builder
	var prevChar rune
	formattedSource := strings.TrimSpace(strings.ToLower(source))

	for _, el := range formattedSource {
		trans, ok := alphabet[el]
		regExp, _ := regexp.Compile("[a-z]")
		match := regExp.MatchString(string(el))

		if el == ' ' {
			if prevChar != '-' {
				result.WriteRune('-')
			}
			prevChar = '-'
		}

		if ok {
			result.WriteString(trans)
			prevChar = rune(trans[len(trans)-1])
		}

		if unicode.IsDigit(el) || match {
			result.WriteRune(el)
			prevChar = el
		}

	}

	return result.String()
}
