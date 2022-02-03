package protovalidator

import "regexp"

const (
	exprEmail           = "^(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	exprHostnameFRC952  = `^[a-zA-Z]([a-zA-Z0-9\-]+[\.]?)*[a-zA-Z0-9]$`                                                                      // https://tools.ietf.org/html/rfc952
	exprHostnameFRC1123 = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?$`                                 // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	exprFQDNFRC1123     = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62})(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*?(\.[a-zA-Z]{1}[a-zA-Z0-9]{0,62})\.?$` // same as exprHostnameFRC1123 but must contain a non numerical TLD (possibly ending with '.')
	exprDataURI         = `^data:((?:\w+\/(?:([^;]|;[^;]).)+)?)`
	exprURLEncoded      = `^(?:[^%]|%[0-9A-Fa-f]{2})*$`
	exprBase64          = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	exprBase64URL       = "^(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})$"
	exprJWT             = "^[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]+\\.[A-Za-z0-9-_]*$"
	exprHexadecimal     = "^(0[xX])?[0-9a-fA-F]+$"
	exprHTMLEncoded     = `&#[x]?([0-9a-fA-F]{2})|(&gt)|(&lt)|(&quot)|(&amp)+[;]?`
	exprHTML            = `<[/]?([a-zA-Z]+).*?>`
)

// Make sure the regex expression is correct.
var (
	regexpEmail           = regexp.MustCompile(exprEmail)
	regexpHostnameFRC952  = regexp.MustCompile(exprHostnameFRC952)
	regexpHostnameFRC1123 = regexp.MustCompile(exprHostnameFRC1123)
	regexpFQDNFRC1123     = regexp.MustCompile(exprFQDNFRC1123)
	regexpDataURI         = regexp.MustCompile(exprDataURI)
	regexpBase64          = regexp.MustCompile(exprBase64)
	regexpBase64URL       = regexp.MustCompile(exprBase64URL)
	regexpURLEncoded      = regexp.MustCompile(exprURLEncoded)
	regexpJWT             = regexp.MustCompile(exprJWT)
	regexpHTML            = regexp.MustCompile(exprHTML)
	regexpHTMLEncoded     = regexp.MustCompile(exprHTMLEncoded)
	regexpHexadecimal     = regexp.MustCompile(exprHexadecimal)
)
