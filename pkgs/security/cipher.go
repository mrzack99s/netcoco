package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"strings"

	"github.com/mrzack99s/netcoco/pkgs/system"
)

func Encrypt(plaintext string) string {

	block, err := aes.NewCipher([]byte(system.SystemConfigVar.MRZ_SW_AUTO.Security.Salt))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	b64CipherText := base64.StdEncoding.EncodeToString(ciphertext)
	b64CipherText = strings.ReplaceAll(b64CipherText, "S", "$")
	b64CipherText = strings.ReplaceAll(b64CipherText, "a", "@")
	b64CipherText = strings.ReplaceAll(b64CipherText, "=", "!")

	return b64CipherText

}

func Decrypt(cipherText []byte) string {

	cipherTextStr := string(cipherText)
	cipherTextStr = strings.ReplaceAll(cipherTextStr, "$", "S")
	cipherTextStr = strings.ReplaceAll(cipherTextStr, "@", "a")
	cipherTextStr = strings.ReplaceAll(cipherTextStr, "!", "=")
	ciphertext, err := base64.StdEncoding.DecodeString(cipherTextStr)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher([]byte(system.SystemConfigVar.MRZ_SW_AUTO.Security.Salt))
	if err != nil {
		panic(err)
	}

	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext)
}
