package urlshortener

import (
	"crypto/sha256"
	"encoding/binary"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func Base62Encode(number uint64) string {
	length := len(alphabet)
	var encodedBuilder strings.Builder
	encodedBuilder.Grow(10)
	for ; number > 0; number = number / uint64(length) {
		encodedBuilder.WriteByte(alphabet[(number % uint64(length))])
	}
	return encodedBuilder.String()
}

func stringToInt64(s string) uint64 {
	hash := sha256.Sum256([]byte(s))
	return binary.BigEndian.Uint64(hash[:8])
}
