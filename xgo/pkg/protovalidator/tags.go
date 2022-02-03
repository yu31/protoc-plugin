package protovalidator

import "fmt"

type TagInfo struct {
	Cond       string      // The if condition.
	Tag        string      // The tag name.
	Value      interface{} // The tag value that user expected.
	FieldValue string      // The variable name of field value.
}

// tag const for oneof
const (
	TagOneOfNotNull = "oneof.not_null"
)

// tag const for float.
const (
	TagFloatEq    = "float.eq"
	TagFloatNe    = "float.ne"
	TagFloatGt    = "float.gt"
	TagFloatLt    = "float.lt"
	TagFloatGte   = "float.gte"
	TagFloatLte   = "float.lte"
	TagFloatIn    = "float.in"
	TagFloatNotIn = "float.not_in"
)

// tag const for int.
const (
	TagIntEq    = "int.eq"
	TagIntNe    = "int.ne"
	TagIntGt    = "int.gt"
	TagIntLt    = "int.lt"
	TagIntGte   = "int.gte"
	TagIntLte   = "int.lte"
	TagIntIn    = "int.in"
	TagIntNotIn = "int.not_in"
)

// tag const for uint.
const (
	TagUintEq    = "uint.eq"
	TagUintNe    = "uint.ne"
	TagUintGt    = "uint.gt"
	TagUintLt    = "uint.lt"
	TagUintGte   = "uint.gte"
	TagUintLte   = "uint.lte"
	TagUintIn    = "uint.in"
	TagUintNotIn = "uint.not_in"
)

// tag const for string.
const (
	TagStringEq    = "string.eq"
	TagStringNe    = "string.ne"
	TagStringGt    = "string.gt"
	TagStringLt    = "string.lt"
	TagStringGte   = "string.gte"
	TagStringLte   = "string.lte"
	TagStringIn    = "string.in"
	TagStringNotIn = "string.not_in"

	TagStringCharLenEq  = "string.char_len_eq"
	TagStringCharLenNe  = "string.char_len_ne"
	TagStringCharLenGt  = "string.char_len_gt"
	TagStringCharLenLt  = "string.char_len_lt"
	TagStringCharLenGte = "string.char_len_gte"
	TagStringCharLenLte = "string.char_len_lte"

	TagStringByteLenEq  = "string.byte_len_eq"
	TagStringByteLenNe  = "string.byte_len_ne"
	TagStringByteLenGt  = "string.byte_len_gt"
	TagStringByteLenLt  = "string.byte_len_lt"
	TagStringByteLenGte = "string.byte_len_gte"
	TagStringByteLenLte = "string.byte_len_lte"

	TagStringRegex = "string.regex"

	TagStringPrefix         = "string.prefix"
	TagStringNoPrefix       = "string.no_prefix"
	TagStringSuffix         = "string.suffix"
	TagStringNoSuffix       = "string.no_suffix"
	TagStringContains       = "string.contains"
	TagStringNotContains    = "string.not_contains"
	TagStringContainsAny    = "string.contains_any"
	TagStringNotContainsAny = "string.not_contains_any"

	TagStringUTF8 = "string.utf8"

	TagStringAscii       = "string.ascii"
	TagStringPrintAscii  = "string.print_ascii"
	TagStringBoolean     = "string.boolean"
	TagStringLowercase   = "string.lowercase"
	TagStringUppercase   = "string.uppercase"
	TagStringAlpha       = "string.alpha"
	TagStringNumber      = "string.number"
	TagStringAlphaNumber = "string.alpha_number"

	TagStringIp              = "string.ip"
	TagStringIpv4            = "string.ipv4"
	TagStringIpv6            = "string.ipv6"
	TagStringIpAddr          = "string.ip_addr"
	TagStringIp4Addr         = "string.ip4_addr"
	TagStringIp6Addr         = "string.ip6_addr"
	TagStringCidr            = "string.cidr"
	TagStringCidrv4          = "string.cidrv4"
	TagStringCidrv6          = "string.cidrv6"
	TagStringMac             = "string.mac"
	TagStringTcpAddr         = "string.tcp_addr"
	TagStringTcp4Addr        = "string.tcp4_addr"
	TagStringTcp6Addr        = "string.tcp6_addr"
	TagStringUdpAddr         = "string.udp_addr"
	TagStringUdp4Addr        = "string.udp4_addr"
	TagStringUdp6Addr        = "string.udp6_addr"
	TagStringUnixAddr        = "string.unix_addr"
	TagStringHostname        = "string.hostname"
	TagStringHostnameRfc1123 = "string.hostname_rfc1123"
	TagStringHostnamePort    = "string.hostname_port"
	TagStringDataURI         = "string.data_uri"
	TagStringFQDN            = "string.fqdn"
	TagStringURI             = "string.uri"
	TagStringURL             = "string.url"
	TagStringURLEncoded      = "string.url_encoded"

	TagStringUnixCron    = "string.unix_cron"
	TagStringEmail       = "string.email"
	TagStringJSON        = "string.json"
	TagStringJWT         = "string.jwt"
	TagStringHTML        = "string.html"
	TagStringHTMLEncoded = "string.html_encoded"
	TagStringBase64      = "string.base64"
	TagStringBase64URL   = "string.base64_url"
	TagStringHexadecimal = "string.hexadecimal"
	TagStringDatetime    = "string.datetime"
	TagStringTimezone    = "string.timezone"
	TagStringUUID        = "string.uuid"
	TagStringUUID1       = "string.uuid1"
	TagStringUUID3       = "string.uuid3"
	TagStringUUID4       = "string.uuid4"
	TagStringUUID5       = "string.uuid5"
)

// tag const for bytes.
const (
	TagBytesLenEq  = "bytes.len_eq"
	TagBytesLenNe  = "bytes.len_ne"
	TagBytesLenGt  = "bytes.len_gt"
	TagBytesLenLt  = "bytes.len_lt"
	TagBytesLenGte = "bytes.len_gte"
	TagBytesLenLte = "bytes.len_lte"
)

// tag const for bool.
const (
	TagBoolEq = "bool.eq"
)

// tag const for enum.
const (
	TagEnumEq      = "enum.eq"
	TagEnumNe      = "enum.ne"
	TagEnumGt      = "enum.gt"
	TagEnumLt      = "enum.lt"
	TagEnumGte     = "enum.gte"
	TagEnumLte     = "enum.lte"
	TagEnumIn      = "enum.in"
	TagEnumNotIn   = "enum.not_in"
	TagEnumInEnums = "enum.in_enums"
)

// tag const for message.
const (
	TagMessageNotNull = "message.not_null"
	TagMessageSkip    = "message.skip"
)

// tag const for repeated.
const (
	TagRepeatedNotNull = "repeated.not_null"
	TagRepeatedLenEq   = "repeated.len_eq"
	TagRepeatedLenNe   = "repeated.len_ne"
	TagRepeatedLenGt   = "repeated.len_gt"
	TagRepeatedLenLt   = "repeated.len_lt"
	TagRepeatedLenGte  = "repeated.len_gte"
	TagRepeatedLenLte  = "repeated.len_lte"
	TagRepeatedUnique  = "repeated.unique"
	TagRepeatedItem    = "repeated.item"
)

// tag const for map.
const (
	TagMapNotNull = "map.not_null"
	TagMapLenEq   = "map.len_eq"
	TagMapLenNe   = "map.len_ne"
	TagMapLenGt   = "map.len_gt"
	TagMapLenLt   = "map.len_lt"
	TagMapLenGte  = "map.len_gte"
	TagMapLenLte  = "map.len_lte"
	TagMapKey     = "map.key"
	TagMapValue   = "map.value"
)

// tagFormatMap is the message format for tag when validate error.
var tagFormatMap = map[string]string{
	// error message for type float.
	TagOneOfNotNull: "the value of %s cannot be null",

	// error message for type float.
	TagFloatEq:    "the value of %s must be equal to '%v'",
	TagFloatNe:    "the value of %s must be not equal to '%v'",
	TagFloatGt:    "the value of %s must be greater than '%v'",
	TagFloatLt:    "the value of %s must be less than '%v'",
	TagFloatGte:   "the value of %s must be greater than or equal to '%v'",
	TagFloatLte:   "the value of %s must be less than or equal to '%v'",
	TagFloatIn:    "the value of %s must be one of in '%v'",
	TagFloatNotIn: "the value of %s must be not one of in '%v'",

	// error message for type int.
	TagIntEq:    "the value of %s must be equal to '%v'",
	TagIntNe:    "the value of %s must be not equal to '%v'",
	TagIntGt:    "the value of %s must be greater than '%v'",
	TagIntLt:    "the value of %s must be less than '%v'",
	TagIntGte:   "the value of %s must be greater than or equal to '%v'",
	TagIntLte:   "the value of %s must be less than or equal to '%v'",
	TagIntIn:    "the value of %s must be one of '%v'",
	TagIntNotIn: "the value of %s must be not one of in '%v'",

	// error message for type uint.
	TagUintEq:    "the value of %s must be equal to '%v'",
	TagUintNe:    "the value of %s must be not equal to '%v'",
	TagUintGt:    "the value of %s must be greater than '%v'",
	TagUintLt:    "the value of %s must be less than '%v'",
	TagUintGte:   "the value of %s must be greater than or equal to '%v'",
	TagUintLte:   "the value of %s must be less than or equal to '%v'",
	TagUintIn:    "the value of %s must be one of '%v'",
	TagUintNotIn: "the value of %s must be not one of in '%v'",

	// error message for type string.
	TagStringEq:    "the value of %s must be equal to '%v'",
	TagStringNe:    "the value of %s must be not equal to '%v'",
	TagStringGt:    "the value of %s must be greater than '%v'",
	TagStringLt:    "the value of %s must be less than '%v'",
	TagStringGte:   "the value of %s must be greater than or equal to '%v'",
	TagStringLte:   "the value of %s must be less than or equal to '%v'",
	TagStringIn:    "the value of %s must be one of '%v'",
	TagStringNotIn: "the value of %s must be not one of in '%v'",

	TagStringCharLenEq:  "the character length of %s must be equal to '%v'",
	TagStringCharLenNe:  "the character length of %s must be not equal to '%v'",
	TagStringCharLenGt:  "the character length of %s must be greater than '%v'",
	TagStringCharLenLt:  "the character length of %s must be less than '%v'",
	TagStringCharLenGte: "the character length of %s must be greater than or equal to '%v'",
	TagStringCharLenLte: "the character length of %s must be less than or equal to '%v'",

	TagStringByteLenEq:  "the byte length of %s must be equal to '%v'",
	TagStringByteLenNe:  "the byte length of %s must be not equal to '%v'",
	TagStringByteLenGt:  "the byte length of %s must be greater than '%v'",
	TagStringByteLenLt:  "the byte length of %s must be less than '%v'",
	TagStringByteLenGte: "the byte length of %s must be greater than or equal to '%v'",
	TagStringByteLenLte: "the byte length of %s must be less than or equal to '%v'",

	TagStringRegex: "the value of %s cannot match regular expression '%s'",

	TagStringPrefix:         "the value of %s must start with string '%v'",
	TagStringNoPrefix:       "the value of %s must not start with string '%v'",
	TagStringSuffix:         "the value of %s must end with string '%v'",
	TagStringNoSuffix:       "the value of %s must not end with string '%v'",
	TagStringContains:       "the value of %s must contains string '%v'",
	TagStringNotContains:    "the value of %s must not contains string '%v'",
	TagStringContainsAny:    "the value of %s must contains any string '%v'",
	TagStringNotContainsAny: "the value of %s must not contains any string '%v'",

	TagStringUTF8:        "the value of %s must be a UTF8 string",
	TagStringAscii:       "the value of %s must be a ASCII characters string",
	TagStringPrintAscii:  "the value of %s must be a printable ASCII string",
	TagStringBoolean:     "the value of %s must be a boolean string",
	TagStringLowercase:   "the value of %s must be a lowercase string",
	TagStringUppercase:   "the value of %s must be a uppercase string",
	TagStringAlpha:       "the value of %s must be a alpha characters string",
	TagStringNumber:      "the value of %s must be a number characters string",
	TagStringAlphaNumber: "the value of %s must be a alphanumeric characters string",

	TagStringIp:              "the value of %s must be a valid IP address",
	TagStringIpv4:            "the value of %s must be a valid IPv4 address",
	TagStringIpv6:            "the value of %s must be a valid IPv6 address",
	TagStringIpAddr:          "the value of %s must be a resolvable IP address",
	TagStringIp4Addr:         "the value of %s must be a resolvable Ipv4 address",
	TagStringIp6Addr:         "the value of %s must be a resolvable Ipv6 address",
	TagStringCidr:            "the value of %s must be a valid CIDR notation",
	TagStringCidrv4:          "the value of %s must be a valid CIDR notation for an IPv4 address",
	TagStringCidrv6:          "the value of %s must be a valid CIDR notation for an IPv6 address",
	TagStringTcpAddr:         "the value of %s must be a valid TCP address",
	TagStringTcp4Addr:        "the value of %s must be a valid IPv4 TCP address",
	TagStringTcp6Addr:        "the value of %s must be a valid IPv6 TCP address",
	TagStringUdpAddr:         "the value of %s must be a valid UDP address",
	TagStringUdp4Addr:        "the value of %s must be a valid UDP v4 address",
	TagStringUdp6Addr:        "the value of %s must be a valid UDP v6 address",
	TagStringMac:             "the value of %s must be a valid MAC address",
	TagStringUnixAddr:        "the value of %s must be a valid UNIX address",
	TagStringHostname:        "the value of %s must be a valid hostname as defined by RFC 952",
	TagStringHostnameRfc1123: "the value of %s must be a valid hostname as defined by RFC 1123",
	TagStringHostnamePort:    "the value of %s must be a string format with hostname and port",
	TagStringDataURI:         "the value of %s must be a string in DATA URI format",
	TagStringFQDN:            "the value of %s must be a string in FQDN format",
	TagStringURI:             "the value of %s must be a string in URI format",
	TagStringURL:             "the value of %s must be a string in URL format",
	TagStringURLEncoded:      "the value of %s must be a string in URL encoded format",

	TagStringUnixCron:    "the value of %s must be a valid standard UNIX-Style crontab expression",
	TagStringEmail:       "the value of %s must be a valid email address as defined by RFC 5322",
	TagStringJSON:        "the value of %s must be a string in JSON format",
	TagStringJWT:         "the value of %s must be a string in JWT format",
	TagStringHTML:        "the value of %s must be a string in HTML format",
	TagStringHTMLEncoded: "the value of %s must be a string in HTML encoded format",
	TagStringBase64:      "the value of %s must be a string in BASE64 format",
	TagStringBase64URL:   "the value of %s must be a string in BASE64 URL format",
	TagStringHexadecimal: "the value of %s must be a string in hexadecimal format",
	TagStringDatetime:    "the value of %s must be a datetime format with layout '%s'",
	TagStringTimezone:    "the value of %s must be a valid timezone",
	TagStringUUID:        "the value of %s must be a valid UUID",
	TagStringUUID1:       "the value of %s must be a valid version 1 UUID",
	TagStringUUID3:       "the value of %s must be a valid version 3 UUID",
	TagStringUUID4:       "the value of %s must be a valid version 4 UUID",
	TagStringUUID5:       "the value of %s must be a valid version 5 UUID",

	// error message for type bytes.
	TagBytesLenEq:  "the length of %s must be equal to '%v'",
	TagBytesLenNe:  "the length of %s must be not equal to '%v'",
	TagBytesLenGt:  "the length of %s must be greater than '%v'",
	TagBytesLenLt:  "the length of %s must be less than '%v'",
	TagBytesLenGte: "the length of %s must be greater than or equal to '%v'",
	TagBytesLenLte: "the length of %s must be less than or equal to '%v'",

	// error message for type bool.
	TagBoolEq: "the value of %s must be equal to '%v'",

	// error message for type enum.
	TagEnumEq:      "the value of %s must be equal to '%v'",
	TagEnumNe:      "the value of %s must be not equal to '%v'",
	TagEnumGt:      "the value of %s must be greater than '%v'",
	TagEnumLt:      "the value of %s must be less than '%v'",
	TagEnumGte:     "the value of %s must be greater than or equal to '%v'",
	TagEnumLte:     "the value of %s must be less than or equal to '%v'",
	TagEnumIn:      "the value of %s must be one of '%v'",
	TagEnumNotIn:   "the value of %s must not be one of '%v'",
	TagEnumInEnums: "the value of %s must in enums of '%v'",

	// error message for type message.
	TagMessageNotNull: "the value of %s cannot be null",

	// error message for type repeated.
	TagRepeatedNotNull: "the value of %s cannot be null",
	TagRepeatedLenEq:   "the length of %s must be equal to '%v'",
	TagRepeatedLenNe:   "the length of %s must be not equal to '%v'",
	TagRepeatedLenGt:   "the length of %s must be greater than '%v'",
	TagRepeatedLenLt:   "the length of %s must be less than '%v'",
	TagRepeatedLenGte:  "the length of %s must be greater than or equal to '%v'",
	TagRepeatedLenLte:  "the length of %s must be less than or equal to '%v'",
	TagRepeatedUnique:  "the array elements in %s must be unique",

	// error message for type map.
	TagMapNotNull: "the value of %s cannot be null",
	TagMapLenEq:   "the length of %s must be equal to '%v'",
	TagMapLenNe:   "the length of %s must be equal not to '%v'",
	TagMapLenGt:   "the length of %s must be greater than '%v'",
	TagMapLenLt:   "the length of %s must be less than '%v'",
	TagMapLenGte:  "the length of %s must be greater than or equal to '%v'",
	TagMapLenLte:  "the length of %s must be less than or equal to '%v'",
}

func BuildErrorReason(tagInfo *TagInfo, fieldDesc string) string {
	format, ok := tagFormatMap[tagInfo.Tag]
	if !ok {
		panic(fmt.Sprintf("message format not defined for tag option %s", tagInfo.Tag))
	}

	var reason string
	if tagInfo.Value == nil {
		reason = fmt.Sprintf(format, fieldDesc)
	} else {
		reason = fmt.Sprintf(format, fieldDesc, tagInfo.Value)
	}
	return reason
}
