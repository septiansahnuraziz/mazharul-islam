package utils

import (
	"bytes"
	"strings"
	"unicode"
)

func SplitString(value, separator string) []string {
	return strings.Split(value, separator)
}

func EscapeQuote(in string) string {
	var buf bytes.Buffer
	const (
		doubleQuoteASCII = 34
		backSlashASCII   = 92
	)

	for i := 0; i < len(in); i++ {
		if in[i] == doubleQuoteASCII && ((i > 0 && in[i-1] != backSlashASCII) || i == 0) {
			buf.WriteByte(backSlashASCII)
		}
		buf.WriteByte(in[i])
	}

	return buf.String()
}

func ToCamelCase(input string) string {
	runes := []rune(input)

	if len(runes) > 0 && unicode.IsUpper(runes[0]) {
		runes[0] = unicode.ToLower(runes[0])
	}

	for i := 1; i < len(runes); i++ {
		if runes[i-1] == ' ' {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}

	return string(runes)
}
