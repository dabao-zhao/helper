package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsFile(t *testing.T) {
	assert.Equal(t, true, IsFile("file.go"))
	assert.Equal(t, false, IsFile("../file"))
}

func TestIsDir(t *testing.T) {
	assert.Equal(t, false, IsDir("file.go"))
	assert.Equal(t, true, IsDir("../file"))
}
