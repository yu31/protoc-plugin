package protovalidator

import (
	"strconv"
	"unicode/utf8"
)

func StringCharsetLenToString(s string) string {
	return strconv.Itoa(utf8.RuneCountInString(s))
}

func StringPointerCharsetLenToString(s *string) string {
	if s == nil {
		return nilStr
	}
	return strconv.Itoa(utf8.RuneCountInString(*s))
}

func StringByteLenToString(s string) string {
	return strconv.Itoa(len(s))
}

func StringPointerByteLenToString(s *string) string {
	if s == nil {
		return nilStr
	}
	return strconv.Itoa(len(*s))
}
