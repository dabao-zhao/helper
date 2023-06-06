package valid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSlice(t *testing.T) {
	assert.Equal(t, true, IsSlice([]int64{1}))
	assert.Equal(t, false, IsSlice([1]int64{1}))
}

func TestIsMap(t *testing.T) {
	assert.Equal(t, true, IsMap(map[string]int64{}))
	assert.Equal(t, false, IsMap([]int64{1}))
}

func TestIsArray(t *testing.T) {
	assert.Equal(t, true, IsArray([1]int64{1}))
	assert.Equal(t, false, IsArray([]int64{1}))
}

func TestIsEmail(t *testing.T) {
	assert.Equal(t, true, IsEmail("1@1.com"))
	assert.Equal(t, false, IsEmail("1@1.c"))
	assert.Equal(t, false, IsEmail("1"))
}

func TestIsBankCardNo(t *testing.T) {
	assert.Equal(t, true, IsBankCardNo("6217003810026896707"))
	assert.Equal(t, false, IsBankCardNo("6217003810026896707111"))
}

func TestIsIDCard(t *testing.T) {
	assert.Equal(t, true, IsIDCard("341226198209165210"))
	assert.Equal(t, false, IsIDCard("941226198209165210"))
}

func TestIsIPv4(t *testing.T) {
	assert.Equal(t, true, IsIPv4("127.0.0.1"))
	assert.Equal(t, false, IsIPv4("127.1222.0.1"))
}

func TestIsIPv6(t *testing.T) {
	assert.Equal(t, true, IsIPv6("CDCD:910A:2222:5498:8475:1111:3900:2020"))
	assert.Equal(t, false, IsIPv6("3900:2020"))
}

func TestIsMobile(t *testing.T) {
	assert.Equal(t, true, IsMobile("19211119993"))
	assert.Equal(t, false, IsMobile("192111199931"))
}
