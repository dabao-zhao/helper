package xmap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var keysTests = []struct {
	s        map[int64]struct{}
	expected []int64
}{
	{
		nil,
		nil,
	},
	{
		map[int64]struct{}{
			1: {},
			2: {},
			3: {},
		},
		[]int64{1, 2, 3},
	},
}

func TestKeys(t *testing.T) {
	for _, test := range keysTests {
		t.Run("", func(t *testing.T) {
			values := Keys(test.s)
			// map 是无序的，要先排序再对比
			sort.Slice(values, func(i, j int) bool {
				return values[i] < values[j]
			})
			assert.Equal(t, test.expected, values)
		})
	}
}

var valuesTests = []struct {
	s        map[int64]int64
	expected []int64
}{
	{
		nil,
		nil,
	},
	{
		map[int64]int64{
			1: 1,
			2: 2,
			3: 3,
		},
		[]int64{1, 2, 3},
	},
}

func TestValues(t *testing.T) {
	for _, test := range valuesTests {
		t.Run("", func(t *testing.T) {
			values := Values(test.s)
			// map 是无序的，要先排序再对比
			sort.Slice(values, func(i, j int) bool {
				return values[i] < values[j]
			})
			assert.Equal(t, test.expected, values)
		})
	}
}
