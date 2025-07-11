# Slugify

A lightweight Go package for converting strings into URL-friendly slugs with Cyrillic support.

## Features

- Converts Cyrillic characters to Latin equivalents
- Handles separators (spaces, dots, dashes, slashes, underscores)
- Supports custom character replacements
- Configurable separator and maximum length
- Automatically removes file extensions
- Clean, predictable output

## Installation

```bash
go get github.com/haadi-coder/slugify
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/haadi-coder/slugify"
)

func main() {
    slug := slugify.Make("Привет мир!")
    fmt.Println(slug) // Output: privet-mir
}
```

### Advanced Usage

```go
options := slugify.Options{
    Separator: "_",
    MaxLength: 20,
    CustomReplacements: map[string]string{
        "&": "and",
        "@": "at",
    },
}

slug := slugify.MakeWithOptions("Тест & проверка", options)
fmt.Println(slug) // Output: test_and_proverka
```

## API

### Functions

- `Make(source string) string` - Convert string to slug with default options
- `MakeWithOptions(inputString string, options Options) string` - Convert with custom options

### Options

```go
type Options struct {
    Separator          string            // Custom separator (default: "-")
    MaxLength          int               // Maximum slug length (0 = unlimited)
    CustomReplacements map[string]string // Custom character replacements
}
```

## Examples

| Input | Output |
|-------|--------|
| `"Привет мир"` | `"privet-mir"` |
| `"Test File.txt"` | `"test-file"` |
| `"Спец/символы_тест"` | `"spets-simvoly-test"` |
| `"Mixed текст 123"` | `"mixed-tekst-123"` |

