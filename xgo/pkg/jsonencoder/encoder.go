package jsonencoder

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"strconv"
)

type Encoder struct {
	escapeHTML bool
	buf        []byte
}

// New return a Encoder.
func New(bufLen int) *Encoder {
	return &Encoder{
		escapeHTML: true,
		buf:        make([]byte, 0, bufLen),
	}
}

// Impls interface io.Writer
func (enc *Encoder) Write(p []byte) (n int, err error) {
	enc.writeBytes(p)
	return len(p), nil
}

// Implements encoder.
func (enc *Encoder) Bytes() []byte {
	return enc.buf
}

func (enc *Encoder) AppendObjectBegin() { enc.writeByte('{') }
func (enc *Encoder) AppendObjectEnd()   { enc.writeByte('}') }

func (enc *Encoder) AppendListBegin() { enc.writeByte('[') }
func (enc *Encoder) AppendListEnd()   { enc.writeByte(']') }

func (enc *Encoder) AppendObjectKey(k string) {
	enc.appendElementSeparator()
	enc.appendString(k)
	enc.writeByte(':')
}

func (enc *Encoder) AppendString(v string) {
	enc.appendElementSeparator()
	enc.appendString(v)
}

func (enc *Encoder) AppendBytes(v []byte) {
	enc.appendElementSeparator()
	if v == nil {
		enc.writeString("null")
		return
	}

	enc.writeByte('"')
	if len(v) != 0 {
		encodedLen := base64.StdEncoding.EncodedLen(len(v))
		// TODO: Improved the alloc logic.
		if encodedLen <= 1024 {
			// The encoded bytes are short enough to allocate for, and
			// Encoding.Encode is still cheaper.
			dst := make([]byte, encodedLen)
			base64.StdEncoding.Encode(dst, v)
			enc.writeBytes(dst)
		} else {
			// The encoded bytes are too long to cheaply allocate, and
			// Encoding.Encode is no longer noticeably cheaper.
			be := base64.NewEncoder(base64.StdEncoding, enc)
			_, _ = be.Write(v)
			_ = be.Close()
		}
	}
	enc.writeByte('"')
}

func (enc *Encoder) AppendBool(v bool) {
	enc.appendElementSeparator()
	enc.buf = strconv.AppendBool(enc.buf, v)
}

func (enc *Encoder) AppendInt32(v int32) {
	enc.appendElementSeparator()
	enc.buf = strconv.AppendInt(enc.buf, int64(v), 10)
}

func (enc *Encoder) AppendInt64(v int64) {
	enc.appendElementSeparator()
	enc.buf = strconv.AppendInt(enc.buf, v, 10)
}

func (enc *Encoder) AppendUint32(v uint32) {
	enc.appendElementSeparator()
	enc.buf = strconv.AppendUint(enc.buf, uint64(v), 10)
}
func (enc *Encoder) AppendUint64(v uint64) {
	enc.appendElementSeparator()
	enc.buf = strconv.AppendUint(enc.buf, v, 10)
}

func (enc *Encoder) AppendFloat32(v float32) {
	enc.appendElementSeparator()
	enc.appendFloat32(v)
}
func (enc *Encoder) AppendFloat64(v float64) {
	enc.appendElementSeparator()
	enc.appendFloat64(v)
}

func (enc *Encoder) AppendNil() {
	enc.appendElementSeparator()
	enc.writeString("null")
}

func (enc *Encoder) AppendInterface(v interface{}) error {
	enc.appendElementSeparator()
	return enc.appendInterface(v)
}

// Add elements separator.
func (enc *Encoder) appendElementSeparator() {
	last := len(enc.buf) - 1
	if last < 0 {
		return
	}

	switch enc.buf[last] {
	case '{', '[', ':', ',':
		return
	default:
		enc.writeByte(',')
	}
}

func (enc *Encoder) appendFloat64(v float64) {
	switch {
	case math.IsNaN(v):
		enc.appendString(`"NaN"`)
	case math.IsInf(v, 1):
		enc.appendString(`"+Inf"`)
	case math.IsInf(v, -1):
		enc.appendString(`"-Inf"`)
	default:
		enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 64)
	}
}

func (enc *Encoder) appendFloat32(x float32) {
	v := float64(x)
	switch {
	case math.IsNaN(v):
		enc.appendString(`"NaN"`)
	case math.IsInf(v, 1):
		enc.appendString(`"+Inf"`)
	case math.IsInf(v, -1):
		enc.appendString(`"-Inf"`)
	default:
		enc.buf = strconv.AppendFloat(enc.buf, v, 'f', -1, 32)
	}
}

func (enc *Encoder) appendInterface(i interface{}) error {
	var err error
	var b []byte

	switch v := i.(type) {
	case nil:
		enc.writeString("null")
		return nil
	case json.Marshaler:
		b, err = v.MarshalJSON()
	default:
		b, err = json.Marshal(i)
		//panic(fmt.Errorf("not support type: %v", v))
	}
	if err != nil {
		return err
	}
	enc.writeBytes(b)
	return nil
}

func (enc *Encoder) writeByte(v byte) {
	enc.buf = append(enc.buf, v)
}

func (enc *Encoder) writeBytes(v []byte) {
	enc.buf = append(enc.buf, v...)
}

func (enc *Encoder) writeString(v string) {
	enc.buf = append(enc.buf, v...)
}
