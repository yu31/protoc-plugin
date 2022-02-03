package protovalidator

import (
	"net"
	"net/url"
	"strconv"
	"strings"
)

// StringIsIP is the validation function for validating if the field's value is a valid v4 or v6 IP address.
func StringIsIP(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil
}

// StringIsIPv4 is the validation function for validating if a value is a valid v4 IP address.
func StringIsIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() != nil
}

// StringIsIPv6 is the validation function for validating if the field's value is a valid v6 IP address.
func StringIsIPv6(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() == nil
}

// StringIsIPAddr is the validation function for validating if the field's value is a resolvable ip address.
func StringIsIPAddr(s string) bool {
	if !StringIsIP(s) {
		return false
	}
	_, err := net.ResolveIPAddr("ip", s)
	return err == nil
}

// StringIsIP4Addr is the validation function for validating if the field's value is a resolvable ip4 address.
func StringIsIP4Addr(s string) bool {
	if !StringIsIPv4(s) {
		return false
	}
	_, err := net.ResolveIPAddr("ip4", s)
	return err == nil
}

// StringIsIP6Addr is the validation function for validating if the field's value is a resolvable ip6 address.
func StringIsIP6Addr(s string) bool {
	if !StringIsIPv6(s) {
		return false
	}
	_, err := net.ResolveIPAddr("ip6", s)
	return err == nil
}

// StringIsCIDR is the validation function for validating if the field's value is a valid v4 or v6 CIDR address.
func StringIsCIDR(s string) bool {
	_, _, err := net.ParseCIDR(s)
	return err == nil
}

// StringIsCIDRv4 is the validation function for validating if the field's value is a valid v4 CIDR address.
func StringIsCIDRv4(s string) bool {
	ip, _, err := net.ParseCIDR(s)
	return err == nil && ip.To4() != nil
}

// StringIsCIDRv6 is the validation function for validating if the field's value is a valid v6 CIDR address.
func StringIsCIDRv6(s string) bool {
	ip, _, err := net.ParseCIDR(s)
	return err == nil && ip.To4() == nil
}

// StringIsMAC is the validation function for validating if the field's value is a valid MAC address.
func StringIsMAC(s string) bool {
	_, err := net.ParseMAC(s)
	return err == nil
}

// StringIsTCPAddr is the validation function for validating if the field's value is a resolvable tcp address.
func StringIsTCPAddr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP4Addr(s) && !StringIsIP6Addr(s) {
	//	return false
	//}
	_, err := net.ResolveTCPAddr("tcp", s)
	return err == nil
}

// StringIsTCP4Addr is the validation function for validating if the field's value is a resolvable tcp4 address.
func StringIsTCP4Addr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP4Addr(s) {
	//	return false
	//}
	_, err := net.ResolveTCPAddr("tcp4", s)
	return err == nil
}

// StringIsTCP6Addr is the validation function for validating if the field's value is a resolvable tcp6 address.
func StringIsTCP6Addr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP6Addr(s) {
	//	return false
	//}
	_, err := net.ResolveTCPAddr("tcp6", s)
	return err == nil
}

// StringIsUDPAddr is the validation function for validating if the field's value is a resolvable udp address.
func StringIsUDPAddr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP4Addr(s) && !StringIsIP6Addr(s) {
	//	return false
	//}
	_, err := net.ResolveUDPAddr("udp", s)
	return err == nil
}

// StringIsUDP4Addr is the validation function for validating if the field's value is a resolvable udp4 address.
func StringIsUDP4Addr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP4Addr(s) {
	//	return false
	//}
	_, err := net.ResolveUDPAddr("udp4", s)
	return err == nil
}

// StringIsUDP6Addr is the validation function for validating if the field's value is a resolvable udp6 address.
func StringIsUDP6Addr(s string) bool {
	if s == "" {
		return false
	}
	//if !StringIsIP6Addr(s) {
	//	return false
	//}
	_, err := net.ResolveUDPAddr("udp6", s)
	return err == nil
}

// StringIsUnixAddr is the validation function for validating if the field's value is a resolvable unix address.
func StringIsUnixAddr(s string) bool {
	if s == "" {
		return false
	}
	_, err := net.ResolveUnixAddr("unix", s)
	return err == nil
}

func StringIsHostname(s string) bool {
	return regexpHostnameFRC952.MatchString(s)
}

func StringIsHostnameRFC1123(s string) bool {
	return regexpHostnameFRC1123.MatchString(s)
}

func StringIsHostnamePort(s string) bool {
	host, port, err := net.SplitHostPort(s)
	if err != nil {
		return false
	}
	// Port must be a iny <= 65535.
	portNum, err := strconv.ParseInt(port, 10, 32)
	if err != nil || portNum > 65535 || portNum < 1 {
		return false
	}

	// If host is specified, it should match a DNS name
	if host != "" {
		return regexpHostnameFRC1123.MatchString(host)
	}
	return true
}

func StringIsFQDN(s string) bool {
	if s == "" {
		return false
	}
	return regexpFQDNFRC1123.MatchString(s)
}

// StringIsDataURI is the validation function for validating if the field's value is a valid data URI.
func StringIsDataURI(s string) bool {
	uri := strings.SplitN(s, ",", 2)
	if len(uri) != 2 {
		return false
	}
	if !regexpDataURI.MatchString(uri[0]) {
		return false
	}
	return regexpBase64.MatchString(uri[1])
}

// StringIsURI is the validation function for validating if the current field's value is a valid URI.
func StringIsURI(s string) bool {
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}
	if len(s) == 0 {
		return false
	}
	_, err := url.ParseRequestURI(s)
	return err == nil
}

// StringIsURL is the validation function for validating if the current field's value is a valid URL.
func StringIsURL(s string) bool {
	// checks needed as of Go 1.6 because of change https://github.com/golang/go/commit/617c93ce740c3c3cc28cdd1a0d712be183d0b328#diff-6c2d018290e298803c0c9419d8739885L195
	// emulate browser and strip the '#' suffix prior to validation. see issue-#237
	if i := strings.Index(s, "#"); i > -1 {
		s = s[:i]
	}
	if len(s) == 0 {
		return false
	}

	_url, err := url.ParseRequestURI(s)
	if err != nil || _url.Scheme == "" {
		return false
	}
	return true
}

func StringIsURLEncoded(s string) bool {
	if s == "" {
		return false
	}
	return regexpURLEncoded.MatchString(s)
}
