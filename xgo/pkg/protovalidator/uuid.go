package protovalidator

import "strings"

func isValidUUIDCharters(s string) bool {
	for i := 0; i < len(s); i++ {
		x := s[i]
		if (x < '0' || x > '9') && (x < 'a' || x > 'f') && (x < 'A' || x > 'F') {
			return false
		}
	}
	return true
}

func stringIsUUID(s string, v byte) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1345][0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	items := strings.Split(s, "-")
	if len(items) != 5 {
		return false
	}
	p1 := items[0]
	if len(p1) != 8 {
		return false
	}
	if !isValidUUIDCharters(p1) {
		return false
	}

	p2 := items[1]
	if len(p2) != 4 {
		return false
	}
	if !isValidUUIDCharters(p2) {
		return false
	}

	p3 := items[2]
	if len(p3) != 4 {
		return false
	}
	if !isValidUUIDCharters(p3) {
		return false
	}

	p4 := items[3]
	if len(p4) != 4 {
		return false
	}
	if !isValidUUIDCharters(p4) {
		return false
	}

	p5 := items[4]
	if !isValidUUIDCharters(p5) {
		return false
	}

	if v == '0' {
		if p3[0] < '1' || p3[0] > '5' {
			return false
		}
	} else {
		if p3[0] != v {
			return false
		}
	}
	return true
}

func StringIsUUID(s string) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1345][0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	return stringIsUUID(s, '0')
}

func StringIsUUID1(s string) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-1[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	return stringIsUUID(s, '1')
}

func StringIsUUID3(s string) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"
	return stringIsUUID(s, '3')
}

func StringIsUUID4(s string) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	return stringIsUUID(s, '4')
}

func StringIsUUID5(s string) bool {
	// "^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"
	return stringIsUUID(s, '5')
}
