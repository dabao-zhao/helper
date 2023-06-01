package xmd5

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	m, err := String("helper")
	assert.Nil(t, err)
	assert.Equal(t, "fde5d67bfb6dc4b598291cc2ce35ee4a", m)
}

func TestFile(t *testing.T) {
	_, err := os.Create("fde5d67bfb6dc4b598291cc2ce35ee4a")
	assert.Nil(t, err)

	m, err := File("fde5d67bfb6dc4b598291cc2ce35ee4a")
	assert.Nil(t, err)
	assert.Equal(t, "\xd4\x1d\x8cُ\x00\xb2\x04\xe9\x80\t\x98\xec\xf8B~", m)

	_ = os.Remove("fde5d67bfb6dc4b598291cc2ce35ee4a")
}
