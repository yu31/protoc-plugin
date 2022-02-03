package govalidator

import (
	"fmt"
	"reflect"

	"github.com/yu31/protoc-plugin/xgo/internal/utils"
	"github.com/yu31/protoc-plugin/xgo/pb/pbvalidator"
	"github.com/yu31/protoc-plugin/xgo/pkg/protovalidator"
	"google.golang.org/protobuf/compiler/protogen"
)

func (p *plugin) loadStringTags(field *protogen.Field, tagOptions *pbvalidator.TagOptions) *pbvalidator.StringTags {
	if tagOptions == nil || tagOptions.Kind == nil {
		return nil
	}

	switch ot := tagOptions.Kind.(type) {
	case *pbvalidator.TagOptions_String_:
		return ot.String_
	default:
		p.exitWithMsg(
			"%s: types <string> only support the kind of TagOptions <string>; and you provided: <%s>",
			p.buildIdentifierWithField(field), reflect.TypeOf(ot).Elem().Name(),
		)
	}
	return nil
}

func (p *plugin) processStringTags(fieldInfo *FieldInfo) []*protovalidator.TagInfo {
	options := p.loadStringTags(fieldInfo.Field, fieldInfo.TagOptions)
	if options == nil {
		return nil
	}

	isPointer := utils.FieldIsPointer(fieldInfo.Field)
	itemName := p.getGoItemName(fieldInfo.Field)

	var tagInfos []*protovalidator.TagInfo
	var cond string

	getFieldValue := func() string {
		var filedValue string
		if isPointer {
			convertMethod := p.g.QualifiedGoIdent(validatorPackage.Ident("StringPointerToString"))
			filedValue = fmt.Sprintf("%s(%s)", convertMethod, itemName)
		} else {
			filedValue = itemName
		}
		return filedValue
	}

	getCharLenFieldValue := func() string {
		var charLenFieldMethod string
		if isPointer {
			charLenFieldMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("StringPointerCharsetLenToString"))
		} else {
			charLenFieldMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("StringCharsetLenToString"))
		}
		charLenFieldValue := fmt.Sprintf("%s(%s)", charLenFieldMethod, itemName)
		return charLenFieldValue
	}

	getByteLenFieldValue := func() string {
		var byteLenFieldMethod string
		if isPointer {
			byteLenFieldMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("StringPointerByteLenToString"))
		} else {
			byteLenFieldMethod = p.g.QualifiedGoIdent(validatorPackage.Ident("StringByteLenToString"))
		}
		byteLenFieldValue := fmt.Sprintf("%s(%s)", byteLenFieldMethod, itemName)
		return byteLenFieldValue
	}

	{
		if options.Eq != nil {
			value := *options.Eq
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s == "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s == "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringEq, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Ne != nil {
			value := *options.Ne
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s != "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s != "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNe, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Gt != nil {
			value := *options.Gt
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s > "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s > "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringGt, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Lt != nil {
			value := *options.Lt
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s < "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s < "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringLt, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Gte != nil {
			value := *options.Gte
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s >= "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s >= "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringGte, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Lte != nil {
			value := *options.Lte
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && *%s <= "%s"`, itemName, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s <= "%s"`, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringLte, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
	}

	{
		if len(options.In) != 0 {
			varName := p.buildVariableNameForTagIn(fieldInfo)
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s[*%s]", itemName, varName, itemName)
			} else {
				cond = fmt.Sprintf("%s[%s]", varName, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIn, Cond: cond, Value: options.In, FieldValue: getFieldValue()})
		}
		if len(options.NotIn) != 0 {
			varName := p.buildVariableNameForTagNotIn(fieldInfo)
			if isPointer {
				cond = fmt.Sprintf("%s != nil && !%s[*%s]", itemName, varName, itemName)
			} else {
				cond = fmt.Sprintf("!%s[%s]", varName, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNotIn, Cond: cond, Value: options.NotIn, FieldValue: getFieldValue()})
		}
	}

	// length of characters.
	{
		if options.CharLenEq != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenEq
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) == %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) == %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenEq, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
		if options.CharLenNe != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenNe
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) != %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) != %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenNe, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
		if options.CharLenGt != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenGt
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) > %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) > %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGt, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
		if options.CharLenLt != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenLt
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) < %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) < %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLt, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
		if options.CharLenGte != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenGte
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) >= %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) >= %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenGte, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
		if options.CharLenLte != nil {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("RuneCountInString"))
			value := *options.CharLenLte
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s) <= %d", itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf("%s(%s) <= %d", method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCharLenLte, Cond: cond, Value: value, FieldValue: getCharLenFieldValue()})
		}
	}

	// length of bytes.
	{
		if options.ByteLenEq != nil {
			value := *options.ByteLenEq
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) == %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) == %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenEq, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
		if options.ByteLenNe != nil {
			value := *options.ByteLenNe
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) != %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) != %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenNe, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
		if options.ByteLenGt != nil {
			value := *options.ByteLenGt
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) > %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) > %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGt, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
		if options.ByteLenLt != nil {
			value := *options.ByteLenLt
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) < %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) < %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLt, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
		if options.ByteLenGte != nil {
			value := *options.ByteLenGte
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) >= %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) >= %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenGte, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
		if options.ByteLenLte != nil {
			value := *options.ByteLenLte
			if isPointer {
				cond = fmt.Sprintf("%s != nil && len(*%s) <= %d", itemName, itemName, value)
			} else {
				cond = fmt.Sprintf("len(%s) <= %d", itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringByteLenLte, Cond: cond, Value: value, FieldValue: getByteLenFieldValue()})
		}
	}

	{
		if options.Regex != nil && *options.Regex != "" {
			varName := p.buildVariableNameForTagRegex(fieldInfo)
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s.MatchString(*%s)", itemName, varName, itemName)
			} else {
				cond = fmt.Sprintf("%s.MatchString(%s)", varName, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringRegex, Cond: cond, Value: *options.Regex, FieldValue: getFieldValue()})
		}

		// string format.
		if options.Prefix != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("HasPrefix"))
			value := *options.Prefix
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && %s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringPrefix, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.NoPrefix != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("HasPrefix"))
			value := *options.NoPrefix
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && !%s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`!%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNoPrefix, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Suffix != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("HasSuffix"))
			value := *options.Suffix
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && %s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringSuffix, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.NoSuffix != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("HasSuffix"))
			value := *options.NoSuffix
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && !%s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`!%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNoSuffix, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Contains != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("Contains"))
			value := *options.Contains
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && %s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringContains, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.NotContains != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("Contains"))
			value := *options.NotContains
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && !%s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`!%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContains, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.ContainsAny != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("ContainsAny"))
			value := *options.ContainsAny
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && %s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringContainsAny, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.NotContainsAny != nil {
			method := p.g.QualifiedGoIdent(stringsPackage.Ident("ContainsAny"))
			value := *options.NotContainsAny
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && !%s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`!%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNotContainsAny, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
	}

	// check codec.
	{
		if options.Utf8 != nil && *options.Utf8 {
			method := p.g.QualifiedGoIdent(utf8Package.Ident("ValidString"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUTF8, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
	}

	// Check charset.
	{
		if options.Ascii != nil && *options.Ascii {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsAscii"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringAscii, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.PrintAscii != nil && *options.PrintAscii {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsPrintAscii"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringPrintAscii, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Boolean != nil && *options.Boolean {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsBoolean"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringBoolean, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Lowercase != nil && *options.Lowercase {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsLowercase"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringLowercase, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uppercase != nil && *options.Uppercase {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUppercase"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUppercase, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Alpha != nil && *options.Alpha {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsAlpha"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringAlpha, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Number != nil && *options.Number {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsNumber"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringNumber, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.AlphaNumber != nil && *options.AlphaNumber {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsAlphaNumber"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringAlphaNumber, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
	}

	// check network.
	{
		if options.Ip != nil && *options.Ip {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIP"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIp, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Ipv4 != nil && *options.Ipv4 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIPv4"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv4, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Ipv6 != nil && *options.Ipv6 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIPv6"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIpv6, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}

		if options.IpAddr != nil && *options.IpAddr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIPAddr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIpAddr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Ip4Addr != nil && *options.Ip4Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIP4Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIp4Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Ip6Addr != nil && *options.Ip6Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsIP6Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringIp6Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Cidr != nil && *options.Cidr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsCIDR"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCidr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Cidrv4 != nil && *options.Cidrv4 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsCIDRv4"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv4, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Cidrv6 != nil && *options.Cidrv6 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsCIDRv6"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringCidrv6, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Mac != nil && *options.Mac {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsMAC"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringMac, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.TcpAddr != nil && *options.TcpAddr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsTCPAddr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringTcpAddr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Tcp4Addr != nil && *options.Tcp4Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsTCP4Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp4Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Tcp6Addr != nil && *options.Tcp6Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsTCP6Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringTcp6Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.UdpAddr != nil && *options.UdpAddr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUDPAddr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUdpAddr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Udp4Addr != nil && *options.Udp4Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUDP4Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp4Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Udp6Addr != nil && *options.Udp6Addr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUDP6Addr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUdp6Addr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.UnixAddr != nil && *options.UnixAddr {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUnixAddr"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixAddr, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Hostname != nil && *options.Hostname {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHostname"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHostname, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.HostnameRfc1123 != nil && *options.HostnameRfc1123 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHostnameRFC1123"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnameRfc1123, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.HostnamePort != nil && *options.HostnamePort {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHostnamePort"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHostnamePort, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.DataUri != nil && *options.DataUri {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsDataURI"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringDataURI, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Fqdn != nil && *options.Fqdn {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsFQDN"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringFQDN, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uri != nil && *options.Uri {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsURI"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringURI, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Url != nil && *options.Url {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsURL"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringURL, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.UrlEncoded != nil && *options.UrlEncoded {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsURLEncoded"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringURLEncoded, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
	}

	// check format.
	{
		if options.UnixCron != nil && *options.UnixCron {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUnixCron"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUnixCron, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Email != nil && *options.Email {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsEmail"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringEmail, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Json != nil && *options.Json {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsJSON"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringJSON, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Jwt != nil && *options.Jwt {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsJWT"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringJWT, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Html != nil && *options.Html {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHTML"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHTML, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.HtmlEncoded != nil && *options.HtmlEncoded {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHTMLEncoded"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHTMLEncoded, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Base64 != nil && *options.Base64 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsBase64"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Base64Url != nil && *options.Base64Url {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsBase64URL"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringBase64URL, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Hexadecimal != nil && *options.Hexadecimal {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsHexadecimal"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringHexadecimal, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Datetime != nil && *options.Datetime != "" {
			value := *options.Datetime
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsDatetime"))
			if isPointer {
				cond = fmt.Sprintf(`%s != nil && %s(*%s, "%s")`, itemName, method, itemName, value)
			} else {
				cond = fmt.Sprintf(`%s(%s, "%s")`, method, itemName, value)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringDatetime, Cond: cond, Value: value, FieldValue: getFieldValue()})
		}
		if options.Timezone != nil && *options.Timezone {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsTimezone"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringTimezone, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uuid != nil && *options.Uuid {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUUID"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uuid1 != nil && *options.Uuid1 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUUID1"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID1, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uuid3 != nil && *options.Uuid3 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUUID3"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID3, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uuid4 != nil && *options.Uuid4 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUUID4"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID4, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
		if options.Uuid5 != nil && *options.Uuid5 {
			method := p.g.QualifiedGoIdent(validatorPackage.Ident("StringIsUUID5"))
			if isPointer {
				cond = fmt.Sprintf("%s != nil && %s(*%s)", itemName, method, itemName)
			} else {
				cond = fmt.Sprintf("%s(%s)", method, itemName)
			}
			tagInfos = append(tagInfos, &protovalidator.TagInfo{Tag: protovalidator.TagStringUUID5, Cond: cond, Value: nil, FieldValue: getFieldValue()})
		}
	}

	return tagInfos
}
