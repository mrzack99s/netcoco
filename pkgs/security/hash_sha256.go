package security

import (
	"crypto/sha256"
	"fmt"
)

func SumSha256(word string) string {
	sum := sha256.Sum256([]byte(word))
	return fmt.Sprintf("%x", sum)
}
