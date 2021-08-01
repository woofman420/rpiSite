package utils

import (
	"encoding/base32"
	"log"
	"os"

	"github.com/cespare/xxhash"
)

// GetCSSHash get hash of the css file.
func GetCSSHash() string {
	newHash := xxhash.New()
	f, err := os.ReadFile("./static/css/main.css")
	if err != nil {
		log.Fatal(err)
	}
	newHash.Write(f)
	return hashForFileName(newHash.Sum(nil))
}

func hashForFileName(hashBytes []byte) string {
	return base32.StdEncoding.EncodeToString(hashBytes)[:8]
}
