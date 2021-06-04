package security

import (
	"crypto/rand"
	"log"
	"math/big"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789$@")

func GeneratePasswordString(length int) (str string) {

	s := make([]rune, length)
	for i := range s {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			log.Fatal(err)
		}
		s[i] = charset[idx.Int64()]
	}
	str = string(s)

	return
}
