package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCases struct {
	name  string
	input string
	want  string
}

type TestCasesWithOptions struct {
	name    string
	input   string
	options Options
	want    string
}

func TestMake(t *testing.T) {
	tests := []TestCases{
		{"Latin letters", "Hello world", "hello-world"},
		{"With special characters", "Hello, world!!!", "hello-world"},
		{"With numbers and special characters", "100% Awesome!!!", "100-awesome"},
		{"In camelCase style", "CamelCase", "camelcase"},
		{"In kebab-case style", "Linux-is-good", "linux-is-good"},

		{"Empty string", "", ""},
		{"String from spaces", "    ", ""},
		{"String from special characters", "!@#$%^&*()", ""},
		{"Sequence of separators", "Hello      World", "hello-world"},
		{"Deleting separators from start and end", "---Hello World---", "hello-world"},

		{"Cyrillic letters", "Москва", "moskva"},
		{"Cyrillic letters with compound vowels_1", "Ёлка", "yolka"},
		{"Cyrillic letters with compound vowels_2", "Щётка", "shchyotka"},
		{"With soft sign))", "семья", "semya"},
		{"With solid sign))", "объект", "obekt"},

		{"URL genereation", "10 советов по Go: как писать идиоматичный код", "10-sovetov-po-go-kak-pisat-idiomatichnyy-kod"},
		{"Filenames", "Отчёт за март 2024.pdf", "otchyot-za-mart-2024"},
		{"Id generation from name", "Смартфон Xiaomi 13 Pro (256ГБ)", "smartfon-xiaomi-13-pro-256gb"},
		{"URL for categories", "Электроника / Телефоны", "elektronika-telefony"},
		{"Tags", "#Новинка!", "novinka"},
	}

	for _, tc := range tests {
		result := Make(tc.input)
		assert.Equal(t, result, tc.want, tc.name)
	}
}

func TestMakeWithOptions(t *testing.T) {

	tests := []TestCasesWithOptions{
		{"Sanke_case separator", "Hello World!", Options{Separator: "_"}, "hello_world"},
		{"Dot separator", "Hello World!", Options{Separator: "."}, "hello.world"},

		{"Custom maxlength limitation", "Очень длинное название", Options{MaxLength: 10}, "ochen-dlin"},

		{"Custom Replacements_1", "Заказ №123", Options{CustomReplacements: map[string]string{"№": "no"}}, "zakaz-no123"},
		{"Custom Replacements_2", "Tom & Jerry", Options{CustomReplacements: map[string]string{"&": "and"}}, "tom-and-jerry"},
		{"Custom Replacements_3", "C++", Options{CustomReplacements: map[string]string{"+": "plus"}}, "c-plus-plus"},
	}

	for _, tc := range tests {
		result := MakeWithOptions(tc.input, tc.options)
		assert.Equal(t, result, tc.want, tc.name)
	}
}
