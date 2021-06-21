package utils

import (
	"encoding/base32"
	"io"

	"github.com/cespare/xxhash"
	"github.com/markbates/pkger"
)

// GetCSSHash get hash of the css file.
func GetCSSHash() string {
	newHash := xxhash.New()
	f, err := pkger.Open("/static/main.css") // can't use a wrapper function or pkger doesn't pack it!
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	newHash.Write(b)
	return hashForFileName(newHash.Sum(nil))
}

func hashForFileName(hashBytes []byte) string {
	return base32.StdEncoding.EncodeToString(hashBytes)[:8]
}
