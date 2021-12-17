package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{-2, 3, 4, 7, 8, 9, 11, 13}

	for idx, num := range arr {
		pos, err := BinarySearch(arr, num)
		assert.Nil(t, err)
		assert.Equal(t, idx, pos)
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	arr := []int{-2, 3, 4, 7, 8, 9, 11, 13}
	target := 11
	_, err := BinarySearch(arr, target)
	if err != nil {
		b.Error(err)
	}
}
