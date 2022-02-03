package jsondecoder

import "unsafe"

// PhasePanicMsg is used as a panic message when we end up with something that
// shouldn't happen. It can indicate a bug in the JSON decoder, or that
// something is editing the data slice while the decoder executes.
const PhasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

type Decoder struct {
	data []byte
	off  int // next read offset in data
	scan scanner

	OpCode OpCode // last read result
}

func New(data []byte) (*Decoder, error) {
	d := &Decoder{
		data: data,
		off:  0,
	}
	err := checkValid(d.data, &d.scan)
	if err != nil {
		return nil, err
	}
	d.scan.reset()
	return d, nil
}

func (d *Decoder) ScanError() error {
	return d.scan.err
}

// ReadIndex readIndex returns the position of the last byte read.
func (d *Decoder) ReadIndex() int {
	return d.off - 1
}

// ReadItem read current item value.
func (d *Decoder) ReadItem() []byte {
	// start index.
	start := d.off - 1
	switch d.OpCode {
	case ScanBeginLiteral:
		d.RescanLiteral()
	case ScanBeginArray, ScanBeginObject:
		d.Skip()
		d.ScanNext()
	default:
		panic(PhasePanicMsg)
	}
	// end index.
	end := d.off - 1
	return d.data[start:end]
}

// ReadObjectKey Read key of object or map.
func (d *Decoder) ReadObjectKey() string {
	item := d.ReadItem()
	key, ok := UnquoteBytes(item)
	if !ok {
		panic(PhasePanicMsg)
	}
	return *(*string)(unsafe.Pointer(&key))
}

func (d *Decoder) ObjectBeforeReadKey() (stop bool) {
	d.ScanWhile(ScanSkipSpace)
	if d.OpCode == ScanEndObject {
		// closing } - can only happen on first iteration.
		return true
	}
	if d.OpCode != ScanBeginLiteral {
		panic(PhasePanicMsg)
	}
	return
}

func (d *Decoder) ObjectBeforeReadValue() {
	if d.OpCode == ScanSkipSpace {
		d.ScanWhile(ScanSkipSpace)
	}
	if d.OpCode != ScanObjectKey {
		panic(PhasePanicMsg)
	}
	d.ScanWhile(ScanSkipSpace)
}

func (d *Decoder) ObjectAfterReadValue() (stop bool) {
	// After read value, Next token must be , or }.
	if d.OpCode == ScanSkipSpace {
		d.ScanWhile(ScanSkipSpace)
	}
	if d.OpCode == ScanEndObject {
		return true
	}
	if d.OpCode != ScanObjectValue {
		panic(PhasePanicMsg)
	}
	return
}

func (d *Decoder) ArrayBeforeReadValue() (stop bool) {
	d.ScanWhile(ScanSkipSpace)
	if d.OpCode == ScanEndArray {
		return true
	}
	return
}

func (d *Decoder) ArrayAfterReadValue() (stop bool) {
	// Next token must be , or ].
	if d.OpCode == ScanSkipSpace {
		d.ScanWhile(ScanSkipSpace)
	}
	if d.OpCode == ScanEndArray {
		return true
	}
	if d.OpCode != ScanArrayValue {
		panic(PhasePanicMsg)
	}
	return
}

// ScanWhile scanWhile processes bytes in d.data[d.off:] until it
// receives a scan code not equal to op.
func (d *Decoder) ScanWhile(op OpCode) {
	s, data, i := &d.scan, d.data, d.off
	for i < len(data) {
		newOp := s.step(s, data[i])
		i++
		if newOp != op {
			d.OpCode = newOp
			d.off = i
			return
		}
	}

	d.off = len(data) + 1 // mark processed EOF with len+1
	d.OpCode = d.scan.eof()
}

// Skip skip scans to the end of what was started.
func (d *Decoder) Skip() {
	s, data, i := &d.scan, d.data, d.off
	depth := len(s.parseState)
	for {
		op := s.step(s, data[i])
		i++
		if len(s.parseState) < depth {
			d.off = i
			d.OpCode = op
			return
		}
	}
}

// ScanNext processes the byte at d.data[d.off].
func (d *Decoder) ScanNext() {
	if d.off < len(d.data) {
		d.OpCode = d.scan.step(&d.scan, d.data[d.off])
		d.off++
	} else {
		d.OpCode = d.scan.eof()
		d.off = len(d.data) + 1 // mark processed EOF with len+1
	}
}

// RescanLiteral is similar to scanWhile(ScanContinue), but it specialises the
// gojson case where we're decoding a literal. The decoder scans the input
// twice, once for syntax errors and to check the length of the value, and the
// second to perform the decoding.
//
// Only in the second step do we use decodeState to tokenize literals, so we
// know there aren't any syntax errors. We can take advantage of that knowledge,
// and scan a literal's bytes much more quickly.
func (d *Decoder) RescanLiteral() {
	data, i := d.data, d.off
Switch:
	switch data[i-1] {
	case '"': // string
		for ; i < len(data); i++ {
			switch data[i] {
			case '\\':
				i++ // escaped char
			case '"':
				i++ // tokenize the closing quote too
				break Switch
			}
		}
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '-': // number
		for ; i < len(data); i++ {
			switch data[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
				'.', 'e', 'E', '+', '-':
			default:
				break Switch
			}
		}
	case 't': // true
		i += len("rue")
	case 'f': // false
		i += len("alse")
	case 'n': // null
		i += len("ull")
	}
	if i < len(data) {
		d.OpCode = stateEndValue(&d.scan, data[i])
	} else {
		d.OpCode = ScanEnd
	}
	d.off = i + 1
}
