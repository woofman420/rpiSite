package utils

import (
	"reflect"
	"unsafe"
)

// This file is pure for functions.
// That mimick originial behaviour, just faster.

func B2s(msg []byte) string {
	return *(*string)(unsafe.Pointer(&msg))
}

// Note it may break if string and/or slice header will change
// in the future go versions.
func S2b(str string) []byte {
	var b []byte
	byteHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	byteHeader.Data = (*reflect.StringHeader)(unsafe.Pointer(&str)).Data

	// This reference is important as without it their is an chance
	// That the str get GC'ed.
	l := len(str)
	byteHeader.Len = l
	byteHeader.Cap = l

	return b
}
