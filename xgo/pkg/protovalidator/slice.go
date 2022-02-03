package protovalidator

import (
	"fmt"
	"reflect"
	"unsafe"
)

func SliceIsUniqueString(a []string) bool {
	cache := make(map[string]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueBytes(a [][]byte) bool {
	cache := make(map[string]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		b := a[i]
		x := *(*string)(unsafe.Pointer(&b))
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueBool(a []bool) bool {
	cache := make(map[bool]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueFloat32(a []float32) bool {
	cache := make(map[float32]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueFloat64(a []float64) bool {
	cache := make(map[float64]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueInt32(a []int32) bool {
	cache := make(map[int32]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueInt64(a []int64) bool {
	cache := make(map[int64]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueUint32(a []uint32) bool {
	cache := make(map[uint32]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueUint64(a []uint64) bool {
	cache := make(map[uint64]struct{}, len(a))
	for i := 0; i < len(a); i++ {
		x := a[i]
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueEnum(a interface{}) bool {
	valueOf := reflect.ValueOf(a)
	if valueOf.Kind() != reflect.Slice {
		panic(fmt.Sprintf("validator: SliceIsUniqueEnum: not support type of v: %s", valueOf.Type().String()))
	}

	valueLen := valueOf.Len()
	cache := make(map[int64]struct{}, valueLen)

	for i := 0; i < valueLen; i++ {
		item := valueOf.Index(i)
		if item.Kind() != reflect.Int32 {
			panic(fmt.Sprintf("validator: SliceIsUniqueEnum: not support type of v: %s", valueOf.Type().String()))
		}
		x := item.Int()
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}

func SliceIsUniqueMessage(a interface{}) bool {
	valueOf := reflect.ValueOf(a)
	if valueOf.Kind() != reflect.Slice {
		panic(fmt.Sprintf("validator: SliceIsUniqueMessage: not support type of v: %s", valueOf.Type().String()))
	}

	valueLen := valueOf.Len()
	cache := make(map[uintptr]struct{}, valueLen)

	for i := 0; i < valueLen; i++ {
		item := valueOf.Index(i)
		x := item.Pointer()
		if _, ok := cache[x]; ok {
			return false
		}
		cache[x] = struct{}{}
	}
	return true
}
