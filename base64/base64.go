package base64

import "encoding/base64"

func DecodeString(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

func EncodeToString(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
