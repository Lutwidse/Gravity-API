package gravity_api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// Author: https://gist.github.com/hothero/7d085573f5cb7cdb5801d7adcf66dcf3

type aesEncrypter struct {
	initialVector string
	passphrase    string
}

func NewAESEncrypter() *aesEncrypter {
	return &aesEncrypter{}
}

func (x *aesEncrypter) SetKeys(k string, iv string) {
	x.initialVector = k
	x.passphrase = iv
}

func (x *aesEncrypter) GetKeys() (string, string) {
	return x.initialVector, x.passphrase
}

func (x aesEncrypter) AESEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(x.initialVector))
	content := []byte(src)
	content = x.PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func (x aesEncrypter) AESDecrypt(crypt []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(x.initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return x.PKCS5Trimming(decrypted)
}

func (x aesEncrypter) PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (aesEncrypter) PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
