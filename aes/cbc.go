package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func CbcEncryptWithIv(key, iv, orig []byte) string {
	block, _ := aes.NewCipher(key)
	plainText := PKCS7Padding(orig, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encrypted := make([]byte, len(plainText))
	blockMode.CryptBlocks(encrypted, plainText)
	return base64.StdEncoding.EncodeToString(encrypted)
}

func CbcDecryptWithIv(key, iv, encrypted []byte) string {
	encrypted, _ = base64.StdEncoding.DecodeString(string(encrypted))
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	orig := make([]byte, len(encrypted))
	blockMode.CryptBlocks(orig, encrypted)
	orig = PKCS7UnPadding(orig)
	return string(orig)
}

func CbcEncrypt(key, orig []byte) string {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	orig = PKCS7Padding(orig, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(orig))
	blockMode.CryptBlocks(encrypted, orig)
	return base64.StdEncoding.EncodeToString(encrypted)
}

func CbcDecrypt(key, encrypted []byte) string {
	encrypted, _ = base64.StdEncoding.DecodeString(string(encrypted))
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	orig := make([]byte, len(encrypted))
	blockMode.CryptBlocks(orig, encrypted)
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
