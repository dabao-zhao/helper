package aes

import (
	"fmt"
	"testing"
)

func TestAesEncryptCBC(t *testing.T) {
	orig := "hello world"
	key := "0123456789012345"
	fmt.Println("origin：", orig)
	encrypted := CbcEncrypt([]byte(key), []byte(orig))
	fmt.Println("encrypt：", encrypted)
	decrypted := CbcDecrypt([]byte(key), []byte(encrypted))
	fmt.Println("decrypt：", decrypted)
}

func TestAesEncryptCBCWithIv(t *testing.T) {
	orig := "hello world"
	key := "0123456789012345"
	iv := "0123456789012345"
	fmt.Println("origin：", orig)
	encrypted := CbcEncryptWithIv([]byte(key), []byte(iv), []byte(orig))
	fmt.Println("encrypt：", encrypted)
	decrypted := CbcDecryptWithIv([]byte(key), []byte(iv), []byte(encrypted))
	fmt.Println("decrypt：", decrypted)
}
