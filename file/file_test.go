package xfile

import (
	"errors"
	"os"
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

func TestFileExists(t *testing.T) {
	assert.Equal(t, false, FileExists("5BaWY92yLs9K4Fvjk5tRbG9TBojbqlnm.go"))
	assert.Equal(t, true, FileExists("file.go"))
}

func TestCreateIfNotExist(t *testing.T) {
	_, err := CreateIfNotExist("file.go")
	assert.Equal(t, errors.New("file.go already exist"), err)

	f, err := CreateIfNotExist("OJliXubJw02Q5kpZKyXoeanynKGlVxSb.go")
	assert.Nil(t, err)
	assert.Equal(t, "OJliXubJw02Q5kpZKyXoeanynKGlVxSb.go", f.Name())

	os.Remove(f.Name())
}

func TestSameFile(t *testing.T) {
	r, err := SameFile("file.go", "file.go")
	assert.Nil(t, err)
	assert.Equal(t, true, r)

	r, err = SameFile("file.go", "file_test.go")
	assert.Nil(t, err)
	assert.Equal(t, false, r)
}
