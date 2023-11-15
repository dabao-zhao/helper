package aes

import "bytes"

func PKCS7UnPadding(text []byte) []byte {
	unPadding := int(text[len(text)-1])
	return text[:(len(text) - unPadding)]
}

func PKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - len(text)%blockSize
	var paddingText []byte
	if padding == 0 {
		paddingText = bytes.Repeat([]byte{byte(blockSize)}, blockSize)
	} else {
		paddingText = bytes.Repeat([]byte{byte(padding)}, padding)
	}
	return append(text, paddingText...)
}
