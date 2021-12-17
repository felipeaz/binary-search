package binary_search

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var arr []int

func init() {
	for i := 0; i < 99999999; i++ {
		arr = append(arr, i)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{-2, 3, 4, 7, 8, 9, 11, 13}
	expectedErr := errors.New("not found")

	for idx, num := range arr {
		pos, err := BinarySearch(arr, num)
		assert.Nil(t, err)
		assert.Equal(t, idx, pos)
	}

	pos, err := BinarySearch(arr, 15)
	assert.Zero(t, pos)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedErr)
}

func TestLinearSearch(t *testing.T) {
	arr := []int{-2, 3, 4, 7, 8, 9, 11, 13}
	expectedErr := errors.New("not found")

	for idx, num := range arr {
		pos, err := LinearSearch(arr, num)
		assert.Nil(t, err)
		assert.Equal(t, idx, pos)
	}

	pos, err := LinearSearch(arr, 15)
	assert.Zero(t, pos)
	assert.NotNil(t, err)
	assert.Equal(t, err, expectedErr)
}

func BenchmarkBinarySearch(b *testing.B) {
	target := 800000
	_, err := BinarySearch(arr, target)
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkLinearSearch(b *testing.B) {
	target := 800000
	_, err := BinarySearch(arr, target)
	if err != nil {
		b.Error(err)
	}
}
