package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"

	"iamricky.com/truck-rental/config"
)

func EncryptPW(plaintext string) (string, error) {
	key := []byte(config.Load("SECRET"))
	iv := generateIV()
	var plainTextBlock []byte
	length := len(plaintext)
	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}
	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)
	str := base64.StdEncoding.EncodeToString(ciphertext)
	return iv + "-" + str, nil
}

func DecryptPW(encrypted string) (string, error) {
	key := []byte(config.Load("SECRET"))
	parts := strings.Split(encrypted, "-")
	iv := ""
	encPW := ""
	if len(parts) >= 2 {
		iv = parts[0]
		encPW = strings.Join(parts[1:], "")
	}
	if iv == "" || encPW == "" {
		return "", errors.New("invalid encrypted string")
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encPW)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("block size cant be zero")
	}
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	length := len(ciphertext)
	unpadding := int(ciphertext[length-1])
	ciphertext = ciphertext[:(length - unpadding)]
	return string(ciphertext), err
}

func generateIV() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
