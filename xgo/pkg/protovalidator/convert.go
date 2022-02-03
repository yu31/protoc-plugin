package protovalidator

import (
	"fmt"
	"reflect"
	"strconv"
)

const nilStr = "<nil>"

func StringPointerToString(v *string) string {
	if v == nil {
		return nilStr
	}
	return *v
}

func Int32ToString(v int32) string {
	return strconv.FormatInt(int64(v), 10)
}

func Int32PointerToString(v *int32) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatInt(int64(*v), 10)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Int64PointerToString(v *int64) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatInt(*v, 10)
}

func Uint32ToString(v uint32) string {
	return strconv.FormatUint(uint64(v), 10)
}

func Uint32PointerToString(v *uint32) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatUint(uint64(*v), 10)
}

func Uint64ToString(v uint64) string {
	return strconv.FormatUint(v, 10)
}

func Uint64PointerToString(v *uint64) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatUint(*v, 10)
}

func Float32ToString(v float32) string {
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func Float32PointerToString(v *float32) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatFloat(float64(*v), 'f', -1, 32)
}

func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func Float64PointerToString(v *float64) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatFloat(*v, 'f', -1, 64)
}

func BoolToString(v bool) string {
	return strconv.FormatBool(v)
}

func BoolPointerToString(v *bool) string {
	if v == nil {
		return nilStr
	}
	return strconv.FormatBool(*v)
}

func EnumPointerToString(v interface{}) string {
	valueOf := reflect.ValueOf(v)
	if valueOf.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("validator: EnumPointerToString: not support type of v: %s", valueOf.Type().String()))
	}
	if valueOf.IsNil() {
		return nilStr
	}
	valueOf = valueOf.Elem()
	if valueOf.Kind() != reflect.Int32 {
		panic(fmt.Sprintf("validator: EnumPointerToString: not support type of v: %s", valueOf.Type().String()))
	}

	return Int64ToString(valueOf.Int())
}
