package xslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var diffTests = []struct {
	s        []int64
	ss       []int64
	expected []int64
}{
	{
		nil,
		nil,
		nil,
	},
	{
		[]int64{1, 2, 3},
		[]int64{2, 3, 4},
		[]int64{1},
	},
	{
		[]int64{2, 3, 4},
		[]int64{1, 2, 3},
		[]int64{4},
	},
}

func TestDiff(t *testing.T) {
	for _, test := range diffTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Diff(test.s, test.ss))
		})
	}
}

var intersectTests = []struct {
	ss       []float64
	params   [][]float64
	expected []float64
}{
	{
		nil,
		nil,
		nil,
	},
	{
		[]float64{1.2, 3.2},
		nil,
		nil,
	},
	{
		nil,
		[][]float64{{1.2, 3.2, 5.5}, {5.5, 1.2}},
		nil,
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2}, {3.2}},
		nil,
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2}},
		[]float64{1.2},
	},
	{
		[]float64{1.2, 3.2},
		[][]float64{{1.2, 3.2, 5.5}, {5.5, 1.2}},
		[]float64{1.2},
	},
}

func TestIntersect(t *testing.T) {
	for _, test := range intersectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Intersect(test.ss, test.params...))
		})
	}
}

var uniqueTests = []struct {
	s        []int64
	expected []int64
}{
	{
		nil,
		nil,
	},
	{
		[]int64{1, 1, 1},
		[]int64{1},
	},
	{
		[]int64{2, 3, 4},
		[]int64{2, 3, 4},
	},
}

func TestUnique(t *testing.T) {
	for _, test := range uniqueTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Unique(test.s))
		})
	}
}

var ContainsTests = []struct {
	s        []int64
	target   int64
	expected bool
}{
	{
		nil,
		0,
		false,
	},
	{
		[]int64{1, 2, 3, 4},
		1,
		true,
	},
	{
		[]int64{2, 3, 4},
		1,
		false,
	},
}

func TestContains(t *testing.T) {
	for _, test := range ContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Contains(test.s, test.target))
		})
	}
}

func TestMerge(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7}, Merge([]int64{1, 2}, []int64{3, 4}, []int64{5, 6, 7}))
}
