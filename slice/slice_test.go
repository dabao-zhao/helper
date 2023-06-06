package xslice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiff(t *testing.T) {
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
	for _, test := range diffTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Diff(test.s, test.ss))
		})
	}
}

func TestIntersect(t *testing.T) {
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
	for _, test := range intersectTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Intersect(test.ss, test.params...))
		})
	}
}

func TestUnique(t *testing.T) {
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

	for _, test := range uniqueTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Unique(test.s))
		})
	}
}

func TestContains(t *testing.T) {
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
	for _, test := range ContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, Contains(test.s, test.target))
		})
	}
}

func TestMerge(t *testing.T) {
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 7}, Merge([]int64{1, 2}, []int64{3, 4}, []int64{5, 6, 7}))
}

func TestFirstOr(t *testing.T) {
	assert.Equal(t, 1, FirstOr([]int{1, 2, 3, 4}, 0))
	assert.Equal(t, 1, FirstOr([]int{}, 1))
}

func TestReverse(t *testing.T) {
	var reverseTests = []struct {
		s        []int64
		expected []int64
	}{
		{
			nil,
			nil,
		},
		{
			[]int64{1, 2, 3},
			[]int64{3, 2, 1},
		},
	}
	for _, test := range reverseTests {
		t.Run("", func(t *testing.T) {
			Reverse(test.s)
			assert.Equal(t, test.expected, test.s)
		})
	}
}

func TestIndexOf(t *testing.T) {
	var indexOfTests = []struct {
		s        []int64
		t        int64
		expected int
	}{
		{
			nil,
			0,
			-1,
		},
		{
			[]int64{1, 2, 3},
			1,
			0,
		},
		{
			[]int64{1, 2, 2, 3},
			2,
			1,
		},
	}
	for _, test := range indexOfTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, IndexOf(test.s, test.t))
		})
	}
}
