package jsonencoder

import (
	"unicode/utf8"
)

const hex = "0123456789abcdef"

// References encoding/json/strings.go: encodeState.string
// NOTE: keep in sync with stringBytes below.
func (enc *Encoder) appendString(s string) {
	enc.writeByte('"')
	start := 0
	for i := 0; i < len(s); {
		if b := s[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!enc.escapeHTML && safeSet[b]) {
				i++
				continue
			}
			if start < i {
				enc.writeString(s[start:i])
			}
			enc.writeByte('\\')
			switch b {
			case '\\', '"':
				enc.writeByte(b)
			case '\n':
				enc.writeByte('n')
			case '\r':
				enc.writeByte('r')
			case '\t':
				enc.writeByte('t')
			default:
				// This encodes bytes < 0x20 except for \t, \n and \r.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				enc.writeString(`u00`)
				enc.writeByte(hex[b>>4])
				enc.writeByte(hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		c, size := utf8.DecodeRuneInString(s[i:])
		if c == utf8.RuneError && size == 1 {
			if start < i {
				enc.writeString(s[start:i])
			}
			enc.writeString(`\ufffd`)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
		if c == '\u2028' || c == '\u2029' {
			if start < i {
				enc.writeString(s[start:i])
			}
			enc.writeString(`\u202`)
			enc.writeByte(hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	if start < len(s) {
		enc.writeString(s[start:])
	}
	enc.writeByte('"')
}

// References encoding/json/strings.go: encodeState.stringBytes
// NOTE: keep in sync with string above.
//func (enc *Encoder) appendBytes(s []byte) {
//	enc.writeByte('"')
//	start := 0
//	for i := 0; i < len(s); {
//		if b := s[i]; b < utf8.RuneSelf {
//			if htmlSafeSet[b] || (!enc.escapeHTML && safeSet[b]) {
//				i++
//				continue
//			}
//			if start < i {
//				enc.writeBytes(s[start:i])
//			}
//			enc.writeByte('\\')
//			switch b {
//			case '\\', '"':
//				enc.writeByte(b)
//			case '\n':
//				enc.writeByte('n')
//			case '\r':
//				enc.writeByte('r')
//			case '\t':
//				enc.writeByte('t')
//			default:
//				// This encodes bytes < 0x20 except for \t, \n and \r.
//				// If escapeHTML is set, it also escapes <, >, and &
//				// because they can lead to security holes when
//				// user-controlled strings are rendered into JSON
//				// and served to some browsers.
//				enc.writeString(`u00`)
//				enc.writeByte(hex[b>>4])
//				enc.writeByte(hex[b&0xF])
//			}
//			i++
//			start = i
//			continue
//		}
//		c, size := utf8.DecodeRune(s[i:])
//		if c == utf8.RuneError && size == 1 {
//			if start < i {
//				enc.writeBytes(s[start:i])
//			}
//			enc.writeString(`\ufffd`)
//			i += size
//			start = i
//			continue
//		}
//		// U+2028 is LINE SEPARATOR.
//		// U+2029 is PARAGRAPH SEPARATOR.
//		// They are both technically valid characters in JSON strings,
//		// but don't work in JSONP, which has to be evaluated as JavaScript,
//		// and can lead to security holes there. It is valid JSON to
//		// escape them, so we do so unconditionally.
//		// See http://timelessrepo.com/json-isnt-a-javascript-subset for discussion.
//		if c == '\u2028' || c == '\u2029' {
//			if start < i {
//				enc.writeBytes(s[start:i])
//			}
//			enc.writeString(`\u202`)
//			enc.writeByte(hex[c&0xF])
//			i += size
//			start = i
//			continue
//		}
//		i += size
//	}
//	if start < len(s) {
//		enc.writeBytes(s[start:])
//	}
//	enc.writeByte('"')
//}
