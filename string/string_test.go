package xstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var toStringTest = []struct {
	s        any
	expected string
}{
	{
		1,
		"1",
	},
	{
		1.1,
		"1.1",
	},
	{
		"123",
		"123",
	},
	{
		'A',
		"65",
	},
	{
		map[string]string{"1": "1"},
		"{\"1\":\"1\"}",
	},
}

func TestToString(t *testing.T) {
	for _, test := range toStringTest {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, ToString(test.s))
		})
	}
}

var case2CamelTest = []struct {
	s        string
	expected string
}{
	{
		"ss_ss",
		"SsSs",
	},
	{
		"",
		"",
	},
}

func TestCase2Camel(t *testing.T) {
	for _, test := range case2CamelTest {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Case2Camel(test.s))
		})
	}
}

var camel2CaseTest = []struct {
	s        string
	expected string
}{
	{
		"SsSs",
		"ss_ss",
	},
	{
		"",
		"",
	},
}

func TestCamel2Case(t *testing.T) {
	for _, test := range camel2CaseTest {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Camel2Case(test.s))
		})
	}
}

func TestUcFirst(t *testing.T) {
	assert.Equal(t, "Qwe", UcFirst("qwe"))
}

func TestLcFirst(t *testing.T) {
	assert.Equal(t, "qwe", LcFirst("Qwe"))
}
