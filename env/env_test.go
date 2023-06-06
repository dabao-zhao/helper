package xenv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOS(t *testing.T) {
	if isw := IsWin(); isw {
		assert.True(t, isw)
		assert.False(t, IsMac())
		assert.False(t, IsLinux())
	}

	if ism := IsMac(); ism {
		assert.True(t, ism)
		assert.False(t, IsWin())
		assert.False(t, IsLinux())
		assert.False(t, IsWSL())
	}

	if isl := IsLinux(); isl {
		assert.True(t, isl)
		assert.False(t, IsMac())
		assert.False(t, IsWin())
		assert.False(t, IsWSL())
	}
}
