package gojson

import (
	"fmt"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (p *plugin) generateUnmarshalCode() {
	msg := p.message
	fields := p.fields

	p.g.P("// UnmarshalJSON for implements json.Unmarshaler.")
	p.g.P("func (this *", msg.GoIdent.GoName, ") UnmarshalJSON(b []byte) error {")
	p.g.P("    if this == nil {")
	p.g.P("        return ", errorsPackage.Ident("New"), "(\"json: Unmarshal: ", string(msg.GoIdent.GoImportPath), ".(*", msg.GoIdent.GoName, ") is nil\")")
	p.g.P("    }")

	if len(fields) >= 0 {
		// Generated flag variables to check oneof.
		for _, field := range fields {
			if utils.FieldIsOneOf(field) {
				options := p.loadOneOfOptions(field.Oneof)
				if !(*options.Ignore) {
					oneOfName := field.Oneof.GoName
					p.g.P("var ", p.genVariableOneofIsStore(oneOfName), " bool")
				}
			}
		}
		// Generated scan code.
		p.unmarshalScanCode()
	}

	p.g.P("    return nil")
	// End function.
	p.g.P("}")
}

func (p *plugin) unmarshalScanCode() {
	p.g.P("")
	p.g.P("decoder, err := ", decoderPackage.Ident("New"), "(b)")
	p.g.P("if err != nil {")
	p.g.P("    return err")
	p.g.P("}")
	p.g.P("")

	p.g.P("// check null.")
	p.g.P("decoder.ScanWhile(", decoderPackage.Ident("ScanSkipSpace"), ")")
	p.g.P("if decoder.OpCode == ", decoderPackage.Ident("ScanBeginLiteral"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	// FIXME: Optimized the error.
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s into object", string(value))`)
	p.g.P("    }")
	p.g.P("    return nil // value is null")
	p.g.P("}")

	p.g.P("if decoder.OpCode != ", decoderPackage.Ident("ScanBeginObject"), " {")
	p.g.P("    panic(", decoderPackage.Ident("PhasePanicMsg"), ")")
	p.g.P("}")

	// generate code for scan object.
	p.g.P("")
	p.g.P("// Scan begin.")
	p.g.P("LOOP_OBJECT:")
	p.g.P("for {")

	p.unmarshalLoopObject()

	// LOOP_OBJECT end.
	p.g.P("}")

	p.g.P("")
	p.g.P("if err = decoder.ScanError(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")
}

func (p *plugin) unmarshalLoopObject() {
	fields := p.fields

	p.g.P("if err = decoder.ScanError(); err != nil {")
	p.g.P("    return err")
	p.g.P("}")

	// Before read key, Read opening " of string key or closing }.
	p.unmarshalObjectBeforeReadKey("LOOP_OBJECT")

	// Read key
	p.g.P("")
	p.g.P("objKey := decoder.ReadObjectKey() // Read key")
	p.g.P("_ = objKey // avoid objKey not used")

	// Before read value
	p.unmarshalObjectBeforeReadValue()

	p.g.P("switch { // process field with key.")

LOOP:
	for _, field := range fields {
		if utils.FieldIsOneOf(field) {
			options := p.loadOneOfOptions(field.Oneof)
			if *options.Ignore {
				continue LOOP
			}
			if *options.HideOneofKey {
				p.unmarshalDecodeOneOf(field.Oneof, "objKey")
				continue LOOP
			}
		}

		var jsonKey string
		if utils.FieldIsOneOf(field) {
			options := p.loadOneOfOptions(field.Oneof)
			if *options.Ignore {
				continue LOOP
			}
			jsonKey = p.getOneOfKey(options, field.Oneof)
		} else {
			options := p.loadFieldOptions(field)
			if *options.Ignore {
				continue LOOP
			}
			jsonKey = p.getFieldKey(options, field)
		}

		p.g.P("case objKey == ", `"`, jsonKey, `"`, ":")
		switch {
		case utils.FieldIsOneOf(field):
			p.unmarshalOneOf(field)
		case field.Desc.IsMap():
			p.unmarshalMap(field)
		case field.Desc.IsList():
			p.unmarshalList(field)
		default:
			p.unmarshalBasic(field)
		}
	}
	p.g.P("default:")
	if *p.msgOptions.DisallowUnknownFields {
		p.g.P("    return ", fmtPackage.Ident("Errorf"), `("json: unknown field %q", objKey)`)
	} else {
		p.g.P("_ = decoder.ReadItem() // discard unknown field")
	}
	// enc switch
	p.g.P("}")

	// After read value..
	p.unmarshalObjectAfterReadValue("LOOP_OBJECT")
}

func (p *plugin) unmarshalOneOf(field *protogen.Field) {
	oneof := field.Oneof

	p.g.P("// decode filed type of oneof;",
		" | field: ", oneof.Desc.FullName(),
		" | GoName: ", oneof.GoName,
	)

	goType := utils.FieldGoType(p.g, field)

	loopLabel := "LOOP_ONEOF_" + p.getOneOfKey(p.loadOneOfOptions(oneof), oneof)

	decodeOneOf := func() {
		//p.g.P("decoder.Skip()")
		p.g.P(loopLabel, ":")
		p.g.P("for {")

		// Before read key, Read opening " of string key or closing }.
		p.unmarshalObjectBeforeReadKey(loopLabel)

		// Read key
		p.g.P("    oneofKey := decoder.ReadObjectKey() // Read key")

		// Before read value
		p.unmarshalObjectBeforeReadValue()

		// Read and process value.
		p.g.P("   switch {")
		p.unmarshalDecodeOneOf(oneof, "oneofKey")
		p.g.P("    default:")
		if *p.msgOptions.DisallowUnknownFields {
			p.g.P("    return ", fmtPackage.Ident("Errorf"), `("json: unknown oneof field %q", oneofKey)`)
		} else {
			p.g.P("_ = decoder.ReadItem() // discard unknown field")
		}
		//p.g.P("       return ", fmtPackage.Ident("Errorf"), `("json: unknown oneof field %q", oneofKey)`)
		// switch end.
		p.g.P("    }")

		p.unmarshalObjectAfterReadValue(loopLabel)
		//p.g.P("    decoder.Skip() ")
		//p.g.P("    break ", loopLabel)

		// end loop.
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", decoderPackage.Ident("ScanBeginLiteral"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as oneof into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    }")
	p.g.P("} else {")
	// check is object.
	p.g.P("    if decoder.OpCode != ", decoderPackage.Ident("ScanBeginObject"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as oneof into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    }")
	decodeOneOf()
	p.g.P("}")
}

func (p *plugin) unmarshalMap(field *protogen.Field) {
	p.g.P("// decode filed type of map;",
		" | field: ", field.Desc.FullName(),
		" | keyKind: ", field.Desc.MapKey().Kind().GoString(),
		" | valueKind: ", field.Desc.MapValue().Kind().GoString(),
		" | goName: ", field.GoName,
	)

	goType := utils.FieldGoType(p.g, field)
	loopLabel := "LOOP_MAP_" + p.getFieldKey(p.loadFieldOptions(field), field)

	decodeKey := func() {
		checkKeyError := func() {
			p.g.P("if err != nil {")
			p.g.P("    return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as map key into field %s of type `, goType, `", key, objKey)`)
			p.g.P("}")
		}
		switch field.Desc.MapKey().Kind() {
		case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
			p.g.P("v, err := ", strconvPackage.Ident("ParseInt"), "(key, 10, 32)")
			checkKeyError()
			p.g.P("mapKey := int32(v)")
		case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
			p.g.P("mapKey, err := ", strconvPackage.Ident("ParseInt"), "(key, 10, 64)")
			checkKeyError()
		case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
			p.g.P("v, err := ", strconvPackage.Ident("ParseUint"), "(key, 10, 32)")
			checkKeyError()
			p.g.P("mapKey := uint32(v)")
		case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
			p.g.P("mapKey, err := ", strconvPackage.Ident("ParseUint"), "(key, 10, 64)")
			checkKeyError()
		case protoreflect.StringKind:
			p.g.P("mapKey := key")
		default:
			panic(fmt.Sprintf(
				"gojson: umarshal: unsupported type of map key, field: %s, kind: %s",
				field.Desc.FullName(), field.Desc.MapKey().Kind().String(),
			))
		}
	}

	decodeMap := func() {
		// create map if not initialized.
		p.g.P("if this.", field.GoName, " == nil { // create map if not initialized.")
		p.g.P("    this.", field.GoName, " = ", "make(", goType, ")")
		p.g.P("}")

		p.g.P(loopLabel, ":")
		p.g.P("for {")

		p.g.P("if err = decoder.ScanError(); err != nil {")
		p.g.P("    return err")
		p.g.P("}")

		// Before read key, Read opening " of string key or closing }.
		p.unmarshalObjectBeforeReadKey(loopLabel)

		// Read key
		p.g.P("")
		p.g.P("key := decoder.ReadObjectKey() // Read map key")

		// Parse key.
		decodeKey()

		// Before read value
		p.unmarshalObjectBeforeReadValue()

		// read and decode value.
		p.unmarshalDecodeValue(field)

		// After read value.
		p.unmarshalObjectAfterReadValue(loopLabel)

		// End loop.
		p.g.P("}")
		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", decoderPackage.Ident("ScanBeginLiteral"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as map into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    } else {")
	p.g.P("        this.", field.GoName, " = nil")
	p.g.P("    }")
	p.g.P("} else {")

	// check is map.
	p.g.P("    if decoder.OpCode != ", decoderPackage.Ident("ScanBeginObject"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as map into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    }")

	decodeMap()

	p.g.P("}")
}

func (p *plugin) unmarshalList(field *protogen.Field) {
	p.g.P("// decode filed type of list;",
		" | field: ", field.Desc.FullName(),
		" | kind: ", field.Desc.Kind().GoString(),
		" | GoName: ", field.GoName,
	)

	goType := utils.FieldGoType(p.g, field)

	loopLabel := "LOOP_LIST_" + p.getFieldKey(p.loadFieldOptions(field), field)

	decodeList := func() {
		p.g.P("if this.", field.GoName, " == nil {")
		p.g.P("    this.", field.GoName, " = ", "make(", goType, ", 0)")
		p.g.P("}")

		p.g.P("i := 0")
		p.g.P("length := len(this.", field.GoName, ")")
		p.g.P(loopLabel, ":")
		p.g.P("for {")

		p.g.P("if err = decoder.ScanError(); err != nil {")
		p.g.P("    return err")
		p.g.P("}")

		// Before Read value.
		p.unmarshalArrayBeforeReadValue(loopLabel)

		// Read list value.
		p.unmarshalDecodeValue(field)

		p.g.P("i++")

		// After read value.
		p.unmarshalArrayAfterReadValue(loopLabel)

		// end LOOP_LIST.
		p.g.P("}")

		// truncate the slice if necessary.
		p.g.P("if i < length {")
		p.g.P("    this.", field.GoName, " = this.", field.GoName, "[:i]")
		p.g.P("}")

		p.g.P("decoder.ScanNext()")
	}

	// Check whether null.
	p.g.P("if decoder.OpCode == ", decoderPackage.Ident("ScanBeginLiteral"), " {")
	p.g.P("    value := decoder.ReadItem()")
	p.g.P("    if value[0] != 'n' {")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as array into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    } else {")
	p.g.P("        this.", field.GoName, " = nil")
	p.g.P("    }")
	p.g.P("} else {")

	// check is list.
	p.g.P("    if decoder.OpCode != ", decoderPackage.Ident("ScanBeginArray"), " {")
	p.g.P("        value := decoder.ReadItem()")
	p.g.P("        return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as array into field %s of type `, goType, `", string(value), objKey)`)
	p.g.P("    }")

	decodeList()

	p.g.P("}")
}

func (p *plugin) unmarshalBasic(field *protogen.Field) {
	if field.Desc.IsMap() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsMap; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsList() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsList; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsWeak() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsWeak; field: %s", field.Desc.FullName()))
	}
	if field.Desc.IsExtension() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsExtension; filed: %s", field.Desc.FullName()))
	}
	if field.Desc.IsPlaceholder() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsPlaceholder; filed: %s", field.Desc.FullName()))
	}
	if field.Desc.IsPacked() {
		panic(fmt.Sprintf("gojson: unmarshal: unsupported case IsPacked; filed: %s", field.Desc.FullName()))
	}

	p.g.P("// decode filed type of basic;",
		" | field: ", field.Desc.FullName(),
		" | kind: ", field.Desc.Kind().GoString(),
		" | GoName: ", field.GoName,
	)

	p.unmarshalDecodeValue(field)
}

func (p *plugin) unmarshalDecodeOneOf(oneof *protogen.Oneof, keyVariable string) {
	for _, field := range oneof.Fields {
		options := p.loadFieldOptions(field)
		if *options.Ignore {
			continue
		}
		p.g.P("case ", keyVariable, " == ", `"`, p.getFieldKey(options, field), `"`, ":")
		p.unmarshalDecodeValue(field)
	}
}

func (p *plugin) unmarshalDecodeValue(field *protogen.Field) {
	options := p.loadFieldOptions(field)

	isMap := field.Desc.IsMap()
	isList := field.Desc.IsList()
	isOneOf := utils.FieldIsOneOf(field)
	isPointer := utils.FieldIsPointer(field)

	goName := field.GoName
	goType := utils.FieldGoType(p.g, field)

	oneOfName := ""
	oneOfType := ""

	if isOneOf {
		oneOfName = field.Oneof.GoName
		oneOfType = p.g.QualifiedGoIdent(field.GoIdent)
	}

	storeValue := func() {
		switch {
		case isMap:
			p.g.P("this.", goName, "[mapKey] = x")
		case isList:
			p.g.P("if i < length {")
			p.g.P("    this.", goName, "[i] = x")
			p.g.P("} else {")
			p.g.P("    this.", goName, " = append(", "this.", goName, ", x", ")")
			p.g.P("}")
		case isOneOf:
			p.g.P("if ", p.genVariableOneofIsStore(oneOfName), " {")
			p.g.P("    return ", fmtPackage.Ident("Errorf"), `("json: unmarshal: the field %s is type oneof, allow contains only one", objKey)`)
			p.g.P("}")
			p.g.P(p.genVariableOneofIsStore(oneOfName), " = true")
			p.g.P("ot := new(", oneOfType, ")")
			p.g.P("ot.", goName, " = x")
			p.g.P("this.", oneOfName, " = ot")
		default:
			if !isPointer {
				p.g.P("this.", goName, " = x")
			} else {
				p.g.P("this.", goName, " = &x")
			}
		}
	}

	returnError := func() {
		switch {
		case isMap:
			p.g.P("return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as map value into field %s of type `, goType, `", string(value), objKey)`)
		case isList:
			p.g.P("return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s as array element into field %s of type `, goType, `", string(value), objKey)`)
		case isOneOf:
			p.g.P("return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s into field %s of type `, goType, `", string(value), objKey)`)
		default:
			p.g.P("return ", fmtPackage.Ident("Errorf"), `("json: cannot unmarshal %s into field %s of type `, goType, `", string(value), objKey)`)
		}
	}
	checkError := func() {
		p.g.P("if err != nil {")
		returnError()
		p.g.P("}")
	}
	checkOk := func() {
		p.g.P("if !ok {")
		returnError()
		p.g.P("}")
	}

	if isMap {
		field = field.Message.Fields[1]
	}

	p.g.P("value := decoder.ReadItem()")

	switch field.Desc.Kind() {
	case protoreflect.DoubleKind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseFloat64"), "(value)")
		checkError()
		storeValue()
	case protoreflect.FloatKind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseFloat32"), "(value)")
		checkError()
		storeValue()
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseInt32"), "(value)")
		checkError()
		storeValue()
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseInt64"), "(value)")
		checkError()
		storeValue()
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseUint32"), "(value)")
		checkError()
		storeValue()
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseUint64"), "(value)")
		checkError()
		storeValue()
	case protoreflect.BoolKind:
		p.g.P("x, err := ", decoderPackage.Ident("ParseBool"), "(value)")
		checkError()
		storeValue()
	case protoreflect.StringKind:
		p.g.P("var x string")
		p.g.P(`if value[0] != 'n' { // 'n' means null`)
		p.g.P("    var ok bool")
		p.g.P("    x, ok = ", decoderPackage.Ident("UnquoteString"), "(value)")
		checkOk()
		p.g.P("}")
		storeValue()
	case protoreflect.BytesKind:
		p.g.P("var x []byte")
		p.g.P("if value[0] != 'n' { // value[0] == 'n' means null")
		p.g.P("    s, ok := ", decoderPackage.Ident("UnquoteBytes"), "(value)")
		checkOk()
		p.g.P("    dst := make([]byte,", base64Package.Ident("StdEncoding"), ".DecodedLen(len(s))", ")")
		p.g.P("    n, err := ", base64Package.Ident("StdEncoding"), ".Decode(dst, s)")
		checkError()
		p.g.P("    x = dst[:n]")
		p.g.P("}")
		storeValue()
	case protoreflect.MessageKind:
		valueType := p.g.QualifiedGoIdent(field.Message.GoIdent)

		p.g.P("var x *", valueType)

		p.g.P("if value[0] != 'n' { // value[0] == 'n' means null")
		switch {
		case isMap:
			p.g.P("x = this.", goName, "[mapKey]")
			p.g.P("if x == nil {")
			p.g.P("    x = new(", valueType, ")")
			p.g.P("}")
		case isList:
			p.g.P("if i < length {")
			p.g.P("    x = this.", goName, "[i]")
			p.g.P("}")
			p.g.P("if x == nil {")
			p.g.P("    x = new(", valueType, ")")
			p.g.P("}")
		case isOneOf:
			p.g.P("    x = new(", valueType, ")")
		default:
			p.g.P("if this.", goName, " == nil {")
			p.g.P("    x = new(", valueType, ")")
			p.g.P("} else {")
			p.g.P("    x = this.", goName)
			p.g.P("}")
		}

		p.g.P("    if um, ok := interface{}(x).(", jsonPackage.Ident("Unmarshaler"), "); ok {")
		p.g.P("        err = um.UnmarshalJSON(value)")
		p.g.P("    } else {")
		p.g.P("        err = ", jsonPackage.Ident("Unmarshal"), "(value, x)")
		p.g.P("    }")
		p.g.P("    if err != nil {")
		p.g.P("        return err")
		p.g.P("    }")
		p.g.P("}")
		storeValue()
	case protoreflect.EnumKind:
		valueType := p.g.QualifiedGoIdent(field.Enum.GoIdent)

		if *options.UseEnumString {
			p.g.P("s, ok := ", decoderPackage.Ident("UnquoteString"), "(value)")
			checkOk()
			p.g.P("x1, ok := ", valueType, "_value[s]")
		} else {
			p.g.P("x1, err := ", decoderPackage.Ident("ParseInt32"), "(value)")
			checkError()
			p.g.P("_, ok := ", valueType, "_name[x1]")
		}

		p.g.P("if !ok {")
		p.g.P("    return ", fmtPackage.Ident("Errorf"), `("json: unknown enum value %s in field %s", string(value), objKey)`)
		p.g.P("}")

		p.g.P("x := ", valueType, "(x1)")
		storeValue()
	default:
		panic(fmt.Sprintf(
			"gojson: unmarshal: unsupported kind of %s as value, field: %s", field.Desc.Kind().String(), field.Desc.FullName(),
		))
	}
}
