package jsondecoder

import (
	"strconv"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

func ParseFloat32(b []byte) (float32, error) {
	v, err := strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func ParseFloat64(b []byte) (float64, error) {
	return strconv.ParseFloat(*(*string)(unsafe.Pointer(&b)), 64)
}

func ParseInt32(b []byte) (int32, error) {
	v, err := strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func ParseInt64(b []byte) (int64, error) {
	return strconv.ParseInt(*(*string)(unsafe.Pointer(&b)), 10, 32)
}

func ParseUint32(b []byte) (uint32, error) {
	v, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func ParseUint64(b []byte) (uint64, error) {
	return strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64)
}

func ParseBool(b []byte) (bool, error) {
	return strconv.ParseBool(*(*string)(unsafe.Pointer(&b)))
}

// UnquoteString converts a quoted JSON string literal s into an actual string t.
// The rules are different than for Go, so cannot use strconv.Unquote.
func UnquoteString(b []byte) (s string, ok bool) {
	b, ok = UnquoteBytes(b)
	if !ok {
		return
	}
	s = *(*string)(unsafe.Pointer(&b))
	return
}

func UnquoteBytes(s []byte) (t []byte, ok bool) {
	if len(s) < 2 || s[0] != '"' || s[len(s)-1] != '"' {
		return
	}
	s = s[1 : len(s)-1]

	// Check for unusual characters. If there are none,
	// then no unquoting is needed, so return a slice of the
	// original bytes.
	r := 0
	for r < len(s) {
		c := s[r]
		if c == '\\' || c == '"' || c < ' ' {
			break
		}
		if c < utf8.RuneSelf {
			r++
			continue
		}
		rr, size := utf8.DecodeRune(s[r:])
		if rr == utf8.RuneError && size == 1 {
			break
		}
		r += size
	}
	if r == len(s) {
		return s, true
	}

	b := make([]byte, len(s)+2*utf8.UTFMax)
	w := copy(b, s[0:r])
	for r < len(s) {
		// Out of room? Can only happen if s is full of
		// malformed UTF-8 and we're replacing each
		// byte with RuneError.
		if w >= len(b)-2*utf8.UTFMax {
			nb := make([]byte, (len(b)+utf8.UTFMax)*2)
			copy(nb, b[0:w])
			b = nb
		}
		switch c := s[r]; {
		case c == '\\':
			r++
			if r >= len(s) {
				return
			}
			switch s[r] {
			default:
				return
			case '"', '\\', '/', '\'':
				b[w] = s[r]
				r++
				w++
			case 'b':
				b[w] = '\b'
				r++
				w++
			case 'f':
				b[w] = '\f'
				r++
				w++
			case 'n':
				b[w] = '\n'
				r++
				w++
			case 'r':
				b[w] = '\r'
				r++
				w++
			case 't':
				b[w] = '\t'
				r++
				w++
			case 'u':
				r--
				rr := getu4(s[r:])
				if rr < 0 {
					return
				}
				r += 6
				if utf16.IsSurrogate(rr) {
					rr1 := getu4(s[r:])
					if dec := utf16.DecodeRune(rr, rr1); dec != unicode.ReplacementChar {
						// A valid pair; consume.
						r += 6
						w += utf8.EncodeRune(b[w:], dec)
						break
					}
					// Invalid surrogate; fall back to replacement rune.
					rr = unicode.ReplacementChar
				}
				w += utf8.EncodeRune(b[w:], rr)
			}

		// Quote, control characters are invalid.
		case c == '"', c < ' ':
			return

		// ASCII
		case c < utf8.RuneSelf:
			b[w] = c
			r++
			w++

		// Coerce to well-formed UTF-8.
		default:
			rr, size := utf8.DecodeRune(s[r:])
			r += size
			w += utf8.EncodeRune(b[w:], rr)
		}
	}
	return b[0:w], true
}
