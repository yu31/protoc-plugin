package protovalidator

import (
	"encoding/json"
	"time"

	"github.com/yu31/cron-go/pkg/expr"
)

func StringIsUnixCron(s string) bool {
	_, err := expr.Standard.Parse(s)
	return err == nil
}

func StringIsEmail(s string) bool {
	return regexpEmail.MatchString(s)
}

func StringIsJSON(s string) bool {
	return json.Valid([]byte(s))
}

func StringIsJWT(s string) bool {
	return regexpJWT.MatchString(s)
}

func StringIsHTML(s string) bool {
	return regexpHTML.MatchString(s)
}

func StringIsHTMLEncoded(s string) bool {
	return regexpHTMLEncoded.MatchString(s)
}

func StringIsBase64(s string) bool {
	return regexpBase64.MatchString(s)
}

func StringIsBase64URL(s string) bool {
	return regexpBase64URL.MatchString(s)
}

func StringIsHexadecimal(s string) bool {
	return regexpHexadecimal.MatchString(s)
}

func StringIsDatetime(s string, layout string) bool {
	_, err := time.Parse(layout, s)
	return err == nil
}

func StringIsTimezone(s string) bool {
	// empty value is converted to UTC by time.LoadLocation but disallow it as it is not a valid time zone name
	if s == "" {
		return false
	}

	//// Local value is converted to the current system time zone by time.LoadLocation but disallow it as it is not a valid time zone name
	//if strings.ToLower(s) == "local" {
	//	return false
	//}

	_, err := time.LoadLocation(s)
	return err == nil
}
