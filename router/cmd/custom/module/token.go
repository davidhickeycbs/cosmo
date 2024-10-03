package module

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
)

func generateHash(desiredLength int, salt, passphrase []byte) []byte {
	data := make([]byte, 0)
	d := make([]byte, 0)

	for len(data) < desiredLength {
		tempD := append(d, passphrase...)
		tempD = append(tempD, salt...)

		hash := md5.Sum(tempD)
		d = hash[:]
		data = append(data, d...)
	}

	return data
}

func Decrypt(token string, key []byte) (string, error) {
	token = replace(token, "-", "/")
	token = replace(token, "_", "+")

	bytes, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(token)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode token: %v", err)
	}

	if len(bytes) < 16 {
		return "", errors.New("invalid token length")
	}

	salted := bytes[:16]
	salt := salted[8:]
	message := bytes[16:]

	hash := generateHash(32+16, salt, key)
	key = hash[:32]
	iv := hash[32:48]

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", fmt.Errorf("could not create cipher: %v", err)
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(message))
	mode.CryptBlocks(decrypted, message)

	// Remove padding
	decrypted = pkcs7Unpad(decrypted)

	return string(decrypted), nil
}

func replace(s, old, new string) string {
	return string(bytes.ReplaceAll([]byte(s), []byte(old), []byte(new)))
}

func pkcs7Unpad(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}
