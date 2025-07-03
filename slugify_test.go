package slugify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Latin letters",
			input: "Hello world",
			want:  "hello-world",
		},
		{
			name:  "With special characters",
			input: "Hello, world!!!",
			want:  "hello-world",
		},
		{
			name:  "With numbers and special characters",
			input: "100% Awesome!!!",
			want:  "100-awesome",
		},
		{
			name:  "In camelCase style",
			input: "CamelCase",
			want:  "camelcase",
		},
		{
			name:  "In kebab-case style",
			input: "Linux-is-good",
			want:  "linux-is-good",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "",
		},
		{
			name:  "String from spaces",
			input: "    ",
			want:  "",
		},
		{
			name:  "String from special characters",
			input: "!@#$%^&*()",
			want:  "",
		},
		{
			name:  "Sequence of separators",
			input: "Hello      World",
			want:  "hello-world",
		},
		{
			name:  "Deleting separators from start and end",
			input: "---Hello World---",
			want:  "hello-world",
		},
		{
			name:  "Dot as separator",
			input: "mail.ru",
			want:  "mail-ru",
		},
		{
			name:  "Cyrillic letters",
			input: "Москва",
			want:  "moskva",
		},
		{
			name:  "Cyrillic letters with compound vowels_1",
			input: "Ёлка",
			want:  "yolka",
		},
		{
			name:  "Cyrillic letters with compound vowels_2",
			input: "Щётка",
			want:  "shchyotka",
		},
		{
			name:  "With soft sign))",
			input: "семья",
			want:  "semya",
		},
		{
			name:  "With solid sign))",
			input: "объект",
			want:  "obekt",
		},
		{
			name:  "URL genereation",
			input: "10 советов по Go: как писать идиоматичный код",
			want:  "10-sovetov-po-go-kak-pisat-idiomatichnyy-kod",
		},
		{
			name:  "Id generation from name",
			input: "Смартфон Xiaomi 13 Pro (256ГБ)",
			want:  "smartfon-xiaomi-13-pro-256gb",
		},
		{
			name:  "URL for categories",
			input: "Электроника / Телефоны",
			want:  "elektronika-telefony",
		},
		{
			name:  "Tags",
			input: "#Новинка!",
			want:  "novinka",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Make(tc.input)
			assert.Equal(t, tc.want, result)
		})
	}

}

func TestMakeWithOptions(t *testing.T) {

	tests := []struct {
		name  string
		input string
		opts  Options
		want  string
	}{
		{
			name:  "Sanke_case separator",
			input: "Hello World!",
			opts:  Options{Separator: "_"},
			want:  "hello_world",
		},
		{
			name:  "Dot separator",
			input: "Hello World!",
			opts:  Options{Separator: "."},
			want:  "hello.world",
		},
		{
			name:  "Custom maxlength limitation",
			input: "Очень длинное название",
			opts:  Options{MaxLength: 10},
			want:  "ochen-dlin",
		},
		{
			name:  "Custom Replacements_1",
			input: "Заказ №123",
			opts:  Options{CustomReplacements: map[rune]string{'№': "no"}},
			want:  "zakaz-no123",
		},
		{
			name:  "Custom Replacements_2",
			input: "Tom & Jerry",
			opts:  Options{CustomReplacements: map[rune]string{'&': "and"}},
			want:  "tom-and-jerry",
		},
		{
			name:  "Custom Replacements_3",
			input: "C++",
			opts:  Options{CustomReplacements: map[rune]string{'+': "plus"}},
			want:  "c-plus-plus",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := MakeWithOptions(tc.input, tc.opts)
			assert.Equal(t, tc.want, result)
		})
	}
}

func TestIdempotency(t *testing.T) {
	t.Run("idempotentcy check", func(t *testing.T) {
		firstResult := Make("10 советов по Go: как писать идиоматичный код")
		secondResult := Make(firstResult)

		assert.Equal(t, secondResult, secondResult)
	})
}
