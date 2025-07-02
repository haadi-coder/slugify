package main

import "testing"

type TestCases struct {
	name  string
	input string
	want  string
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
		{"Cyrillic letters with compound vowels", "Ёлка", "yolka"},
		{"Cyrillic letters with compound vowels", "Щётка", "shchyotka"},
		{"With soft sign))", "семья", "semya"},
		{"With solid sign))", "объект", "obekt"},

		{"URL genereation", "10 советов по Go: как писать идиоматичный код", "10-sovetov-po-go-kak-pisat-idiomatichnyy-kod"},
		{"Filenames", "Отчёт за март 2024.pdf", "otchyot-za-mart-2024"},
		{"Id generation from name", "Смартфон Xiaomi 13 Pro (256ГБ)", "smartfon-xiaomi-13-pro-256gb"},
		{"URL for categories", "Электроника / Телефоны", "elektronika-telefony"},
		{"Tags", "#Новинка!", "novinka"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Make(tc.input)

			if result != tc.want {
				t.Errorf("Result was incorrect, got: %s, want: %s.", result, tc.want)
			}
		})
	}
}
