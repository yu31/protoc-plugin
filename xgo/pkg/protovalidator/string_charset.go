package protovalidator

import (
	"strconv"
	"strings"
	"unicode"
)

// StringIsAscii check whether the string is ascii charset.
func StringIsAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func StringIsPrintAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < 33 {
			return false
		}
	}
	return true
}

func StringIsBoolean(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

func StringIsLowercase(s string) bool {
	return s == strings.ToLower(s)
}

func StringIsUppercase(s string) bool {
	return s == strings.ToUpper(s)
}

// StringIsAlpha allowed a ~ z and A ~ Z
func StringIsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		x := s[i]
		if (x < 'a' || x > 'z') && (x < 'A' || x > 'Z') {
			return false
		}
	}
	return true
}

func StringIsNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		x := s[i]
		if x < '0' || x > '9' {
			return false
		}
	}
	return true
}

func StringIsAlphaNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		x := s[i]
		if (x < 'a' || x > 'z') && (x < 'A' || x > 'Z') && (x < '0' || x > '9') {
			return false
		}
	}
	return true
}
