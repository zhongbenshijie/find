package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	intSlice := []int{1, 2, 5, 6, 29, 58, 189, 401}
	target := 5

	position := BinarySearch(intSlice, target)
	if -1 == position {
		t.Log("target not in intSlice")
	} else {
		t.Log("target finded, position : " + fmt.Sprint(position))
	}
}
