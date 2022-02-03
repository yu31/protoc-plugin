package jsondecoder

// JSON value parser state machine.
// Just about at the limit of what is reasonable to write by hand.
// Some parts are a bit tedious, but overall it nicely factors out the
// otherwise gojsonexternal code from the multiple scanning functions
// in this package (Compact, Indent, checkValid, etc).
//
// This file starts with two simple examples using the scanner
// before diving into the scanner itself.

// checkValid verifies that data is valid JSON-encoded data.
// scan is passed in for use by checkValid to avoid an allocation.
func checkValid(data []byte, scan *scanner) error {
	scan.reset()
	for _, c := range data {
		scan.bytes++
		if scan.step(scan, c) == ScanError {
			return scan.err
		}
	}
	if scan.eof() == ScanError {
		return scan.err
	}
	return nil
}

type OpCode int8

// These values are returned by the state transition functions
// assigned to scanner.state and the method scanner.eof.
// They give details about the current state of the scan that
// callers might be interested to know about.
// It is okay to ignore the return value of any particular
// call to scanner.state: if one call returns ScanError,
// every subsequent call will return ScanError too.
const (
	// Continue.
	ScanContinue     OpCode = iota // uninteresting byte
	ScanBeginLiteral               // end implied by next result != ScanContinue
	ScanBeginObject                // begin object
	ScanObjectKey                  // just finished object key (string)
	ScanObjectValue                // just finished non-last object value
	ScanEndObject                  // end object (implies ScanObjectValue if possible)
	ScanBeginArray                 // begin array
	ScanArrayValue                 // just finished array value
	ScanEndArray                   // end array (implies ScanArrayValue if possible)
	ScanSkipSpace                  // space byte; can skip; known to be last "continue" result

	// Stop.
	ScanEnd   // top-level value ended *before* this byte; known to be first "stop" result
	ScanError // hit an error, scanner.err.
)

type ParsePhase int8

// These values are stored in the parseState stack.
// They give the current state of a composite value
// being scanned. If the parser is inside a nested value
// the parseState describes the nested state, outermost at entry 0.
const (
	parseObjectKey   ParsePhase = iota // parsing object key (before colon)
	parseObjectValue                   // parsing object value (after colon)
	parseArrayValue                    // parsing array value
)

// This limits the max nesting depth to prevent stack overflow.
// This is permitted by https://tools.ietf.org/html/rfc7159#section-9
const maxNestingDepth = 10000

// A scanner is a JSON scanning state machine.
// Callers call scan.reset and then pass bytes in one at a time
// by calling scan.step(&scan, c) for each byte.
// The return value, referred to as an opcode, tells the
// caller about significant parsing events like beginning
// and ending literals, objects, and arrays, so that the
// caller can follow along if it wishes.
// The return value ScanEnd indicates that a single top-level
// JSON value has been completed, *before* the byte that
// just got passed in.  (The indication must be delayed in order
// to recognize the end of numbers: is 123 a whole value or
// the beginning of 12345e+6?).
type scanner struct {
	// The step is a func to be called to execute the next transition.
	// Also tried using an integer constant and a single func
	// with a switch, but using the func directly was 10% faster
	// on a 64-bit Mac Mini, and it's nicer to read.
	step func(*scanner, byte) OpCode

	// Reached end of top-level value.
	endTop bool

	// Stack of what we're in the middle of - array values, object keys, object values.
	parseState []ParsePhase

	// Error that happened, if any.
	err error

	// total bytes consumed, updated by decoder.Decode (and deliberately
	// not set to zero by scan.reset)
	bytes int64
}

// reset prepares the scanner for use.
// It must be called before calling s.step.
func (s *scanner) reset() {
	s.step = stateBeginValue
	s.parseState = s.parseState[0:0]
	s.err = nil
	s.endTop = false
}

// eof tells the scanner that the end of input has been reached.
// It returns a scan status just as s.step does.
func (s *scanner) eof() OpCode {
	if s.err != nil {
		return ScanError
	}
	if s.endTop {
		return ScanEnd
	}
	s.step(s, ' ')
	if s.endTop {
		return ScanEnd
	}
	if s.err == nil {
		s.err = &SyntaxError{"unexpected end of JSON input", s.bytes}
	}
	return ScanError
}

// pushParseState pushes a new parse state p onto the parse stack.
// an error state is returned if maxNestingDepth was exceeded, otherwise successState is returned.
func (s *scanner) pushParseState(c byte, parseState ParsePhase, successState OpCode) OpCode {
	s.parseState = append(s.parseState, parseState)
	if len(s.parseState) <= maxNestingDepth {
		return successState
	}
	return s.error(c, "exceeded max depth")
}

// popParseState pops a parse state (already obtained) off the stack
// and updates s.step accordingly.
func (s *scanner) popParseState() {
	n := len(s.parseState) - 1
	s.parseState = s.parseState[0:n]
	if n == 0 {
		s.step = stateEndTop
		s.endTop = true
	} else {
		s.step = stateEndValue
	}
}

// stateBeginValueOrEmpty is the state after reading `[`.
func stateBeginValueOrEmpty(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return ScanSkipSpace
	}
	if c == ']' {
		return stateEndValue(s, c)
	}
	return stateBeginValue(s, c)
}

// stateBeginValue is the state at the beginning of the input.
func stateBeginValue(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return ScanSkipSpace
	}
	switch c {
	case '{':
		s.step = stateBeginStringOrEmpty
		return s.pushParseState(c, parseObjectKey, ScanBeginObject)
	case '[':
		s.step = stateBeginValueOrEmpty
		return s.pushParseState(c, parseArrayValue, ScanBeginArray)
	case '"':
		s.step = stateInString
		return ScanBeginLiteral
	case '-':
		s.step = stateNeg
		return ScanBeginLiteral
	case '0': // beginning of 0.123
		s.step = state0
		return ScanBeginLiteral
	case 't': // beginning of true
		s.step = stateT
		return ScanBeginLiteral
	case 'f': // beginning of false
		s.step = stateF
		return ScanBeginLiteral
	case 'n': // beginning of null
		s.step = stateN
		return ScanBeginLiteral
	}
	if '1' <= c && c <= '9' { // beginning of 1234.5
		s.step = state1
		return ScanBeginLiteral
	}
	return s.error(c, "looking for beginning of value")
}

// stateBeginStringOrEmpty is the state after reading `{`.
func stateBeginStringOrEmpty(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return ScanSkipSpace
	}
	if c == '}' {
		n := len(s.parseState)
		s.parseState[n-1] = parseObjectValue
		return stateEndValue(s, c)
	}
	return stateBeginString(s, c)
}

// stateBeginString is the state after reading `{"key": value,`.
func stateBeginString(s *scanner, c byte) OpCode {
	if isSpace(c) {
		return ScanSkipSpace
	}
	if c == '"' {
		s.step = stateInString
		return ScanBeginLiteral
	}
	return s.error(c, "looking for beginning of object key string")
}

// stateEndValue is the state after completing a value,
// such as after reading `{}` or `true` or `["x"`.
func stateEndValue(s *scanner, c byte) OpCode {
	n := len(s.parseState)
	if n == 0 {
		// Completed top-level before the current byte.
		s.step = stateEndTop
		s.endTop = true
		return stateEndTop(s, c)
	}
	if isSpace(c) {
		s.step = stateEndValue
		return ScanSkipSpace
	}
	ps := s.parseState[n-1]
	switch ps {
	case parseObjectKey:
		if c == ':' {
			s.parseState[n-1] = parseObjectValue
			s.step = stateBeginValue
			return ScanObjectKey
		}
		return s.error(c, "after object key")
	case parseObjectValue:
		if c == ',' {
			s.parseState[n-1] = parseObjectKey
			s.step = stateBeginString
			return ScanObjectValue
		}
		if c == '}' {
			s.popParseState()
			return ScanEndObject
		}
		return s.error(c, "after object key:value pair")
	case parseArrayValue:
		if c == ',' {
			s.step = stateBeginValue
			return ScanArrayValue
		}
		if c == ']' {
			s.popParseState()
			return ScanEndArray
		}
		return s.error(c, "after array element")
	}
	return s.error(c, "")
}

// stateEndTop is the state after finishing the top-level value,
// such as after reading `{}` or `[1,2,3]`.
// Only space characters should be seen now.
func stateEndTop(s *scanner, c byte) OpCode {
	if !isSpace(c) {
		// Complain about non-space byte on next call.
		s.error(c, "after top-level value")
	}
	return ScanEnd
}

// stateInString is the state after reading `"`.
func stateInString(s *scanner, c byte) OpCode {
	if c == '"' {
		s.step = stateEndValue
		return ScanContinue
	}
	if c == '\\' {
		s.step = stateInStringEsc
		return ScanContinue
	}
	if c < 0x20 {
		return s.error(c, "in string literal")
	}
	return ScanContinue
}

// stateInStringEsc is the state after reading `"\` during a quoted string.
func stateInStringEsc(s *scanner, c byte) OpCode {
	switch c {
	case 'b', 'f', 'n', 'r', 't', '\\', '/', '"':
		s.step = stateInString
		return ScanContinue
	case 'u':
		s.step = stateInStringEscU
		return ScanContinue
	}
	return s.error(c, "in string escape code")
}

// stateInStringEscU is the state after reading `"\u` during a quoted string.
func stateInStringEscU(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU1
		return ScanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU1 is the state after reading `"\u1` during a quoted string.
func stateInStringEscU1(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU12
		return ScanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU12 is the state after reading `"\u12` during a quoted string.
func stateInStringEscU12(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInStringEscU123
		return ScanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateInStringEscU123 is the state after reading `"\u123` during a quoted string.
func stateInStringEscU123(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F' {
		s.step = stateInString
		return ScanContinue
	}
	// numbers
	return s.error(c, "in \\u hexadecimal character escape")
}

// stateNeg is the state after reading `-` during a number.
func stateNeg(s *scanner, c byte) OpCode {
	if c == '0' {
		s.step = state0
		return ScanContinue
	}
	if '1' <= c && c <= '9' {
		s.step = state1
		return ScanContinue
	}
	return s.error(c, "in numeric literal")
}

// state1 is the state after reading a non-zero integer during a number,
// such as after reading `1` or `100` but not `0`.
func state1(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = state1
		return ScanContinue
	}
	return state0(s, c)
}

// state0 is the state after reading `0` during a number.
func state0(s *scanner, c byte) OpCode {
	if c == '.' {
		s.step = stateDot
		return ScanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return ScanContinue
	}
	return stateEndValue(s, c)
}

// stateDot is the state after reading the integer and decimal point in a number,
// such as after reading `1.`.
func stateDot(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = stateDot0
		return ScanContinue
	}
	return s.error(c, "after decimal point in numeric literal")
}

// stateDot0 is the state after reading the integer, decimal point, and subsequent
// digits of a number, such as after reading `3.14`.
func stateDot0(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		return ScanContinue
	}
	if c == 'e' || c == 'E' {
		s.step = stateE
		return ScanContinue
	}
	return stateEndValue(s, c)
}

// stateE is the state after reading the mantissa and e in a number,
// such as after reading `314e` or `0.314e`.
func stateE(s *scanner, c byte) OpCode {
	if c == '+' || c == '-' {
		s.step = stateESign
		return ScanContinue
	}
	return stateESign(s, c)
}

// stateESign is the state after reading the mantissa, e, and sign in a number,
// such as after reading `314e-` or `0.314e+`.
func stateESign(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		s.step = stateE0
		return ScanContinue
	}
	return s.error(c, "in exponent of numeric literal")
}

// stateE0 is the state after reading the mantissa, e, optional sign,
// and at least one digit of the exponent in a number,
// such as after reading `314e-2` or `0.314e+1` or `3.14e0`.
func stateE0(s *scanner, c byte) OpCode {
	if '0' <= c && c <= '9' {
		return ScanContinue
	}
	return stateEndValue(s, c)
}

// stateT is the state after reading `t`.
func stateT(s *scanner, c byte) OpCode {
	if c == 'r' {
		s.step = stateTr
		return ScanContinue
	}
	return s.error(c, "in literal true (expecting 'r')")
}

// stateTr is the state after reading `tr`.
func stateTr(s *scanner, c byte) OpCode {
	if c == 'u' {
		s.step = stateTru
		return ScanContinue
	}
	return s.error(c, "in literal true (expecting 'u')")
}

// stateTru is the state after reading `tru`.
func stateTru(s *scanner, c byte) OpCode {
	if c == 'e' {
		s.step = stateEndValue
		return ScanContinue
	}
	return s.error(c, "in literal true (expecting 'e')")
}

// stateF is the state after reading `f`.
func stateF(s *scanner, c byte) OpCode {
	if c == 'a' {
		s.step = stateFa
		return ScanContinue
	}
	return s.error(c, "in literal false (expecting 'a')")
}

// stateFa is the state after reading `fa`.
func stateFa(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateFal
		return ScanContinue
	}
	return s.error(c, "in literal false (expecting 'l')")
}

// stateFal is the state after reading `fal`.
func stateFal(s *scanner, c byte) OpCode {
	if c == 's' {
		s.step = stateFals
		return ScanContinue
	}
	return s.error(c, "in literal false (expecting 's')")
}

// stateFals is the state after reading `fals`.
func stateFals(s *scanner, c byte) OpCode {
	if c == 'e' {
		s.step = stateEndValue
		return ScanContinue
	}
	return s.error(c, "in literal false (expecting 'e')")
}

// stateN is the state after reading `n`.
func stateN(s *scanner, c byte) OpCode {
	if c == 'u' {
		s.step = stateNu
		return ScanContinue
	}
	return s.error(c, "in literal null (expecting 'u')")
}

// stateNu is the state after reading `nu`.
func stateNu(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateNul
		return ScanContinue
	}
	return s.error(c, "in literal null (expecting 'l')")
}

// stateNul is the state after reading `nul`.
func stateNul(s *scanner, c byte) OpCode {
	if c == 'l' {
		s.step = stateEndValue
		return ScanContinue
	}
	return s.error(c, "in literal null (expecting 'l')")
}

// stateError is the state after reaching a syntax error,
// such as after reading `[1}` or `5.1.2`.
func stateError(s *scanner, c byte) OpCode {
	return ScanError
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) OpCode {
	s.step = stateError
	s.err = &SyntaxError{"invalid character " + quoteChar(c) + " " + context, s.bytes}
	return ScanError
}
