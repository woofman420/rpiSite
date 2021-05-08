package utils

import (
	"encoding/base64"
	"reflect"
	"unsafe"
)

// This file is pure for functions.
// That mimick originial behaviour, just faster.

// UnsafeStringConversion is a quick dirty way to avoid allocations from byte to string.
func UnsafeStringConversion(msg []byte) string {
	return *(*string)(unsafe.Pointer(&msg))
}

// UnsafeByteConversion will may break if string and/or slice header will change
// in the future go versions.
func UnsafeByteConversion(str string) []byte {
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

// EncodeToString will Base64(URL Variant) encode the given []byte.
func EncodeToString(src []byte) string {
	buf := make([]byte, (len(src)*8+5)/6)
	base64.RawURLEncoding.Encode(buf, src)
	return UnsafeStringConversion(buf)
}

// DecodeString will Base64(URL Variant) decode the given string.
func DecodeString(s string) ([]byte, error) {
	dbuf := make([]byte, len(s)*6/8)
	n, err := base64.RawURLEncoding.Decode(dbuf, UnsafeByteConversion(s))
	return dbuf[:n], err
}
