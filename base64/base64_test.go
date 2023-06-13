package base64

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeToString(t *testing.T) {
	s := EncodeToString([]byte("encode"))
	assert.Equal(t, "ZW5jb2Rl", s)
}

func TestDecodeString(t *testing.T) {
	b, err := DecodeString("ZW5jb2Rl")
	assert.Nil(t, err)
	assert.Equal(t, "encode", string(b))
}
