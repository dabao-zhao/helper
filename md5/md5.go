package xmd5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

func String(str string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil)), err
}

func File(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return string(h.Sum(nil)), nil
}
