package protovalidator

import (
	"strings"
)

const errorId = "ValidateError"
const withValueMsg = " and you provide "

type ValidateError struct {
	//// The validator tag name.
	//Tag string
	//
	//// The value that you expected.
	//ExpectedValue string
	//
	//// The field's value will return in message. It may be "" when no need display.
	//FieldValue string

	message string
}

func (e *ValidateError) Error() string {
	return e.message
}

func FieldError1(structName string, reason string, value string) error {
	//message := fmt.Sprintf("ValidateError: <%s>: %s and you provide '%+v'", structName, reason, value)

	strLen := len(errorId) + len(structName) + 6 + len(reason) + len(withValueMsg) + 2 + len(value)

	var s strings.Builder
	s.Grow(strLen)

	s.WriteString(errorId)
	s.WriteString(": ")
	s.WriteString("<")
	s.WriteString(structName)
	s.WriteString(">")
	s.WriteString(": ")
	s.WriteString(reason)
	s.WriteString(withValueMsg)
	s.WriteString("'")
	s.WriteString(value)
	s.WriteString("'")

	message := s.String()

	e := &ValidateError{
		message: message,
	}
	return e
}

func FieldError2(structName string, reason string) error {
	//message := fmt.Sprintf("ValidateError: <%s>: %s", structName, reason)

	strLen := len(errorId) + len(structName) + 6 + len(reason)

	var s strings.Builder
	s.Grow(strLen)

	s.WriteString(errorId)
	s.WriteString(": ")
	s.WriteString("<")
	s.WriteString(structName)
	s.WriteString(">")
	s.WriteString(": ")
	s.WriteString(reason)

	message := s.String()
	e := &ValidateError{
		message: message,
	}
	return e
}
