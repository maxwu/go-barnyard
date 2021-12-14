package binsearch

import (
	"testing"
)

func TestBinSearch(t *testing.T) {
	sortedSlice := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}

	p := BinSearch(sortedSlice, 11)
	if p != 7 {
		t.Errorf("Expected %d in position 7 but returned %d", 11, p)
	}

	sortedSlice = []int{1, 3, 4, 6, 7, 9, 10, 11, 13}

	p = BinSearch(sortedSlice, 3)
	if p != 1 {
		t.Errorf("Expected %d in position 1 but returned %d", 3, p)
	}

	p = BinSearch(sortedSlice, 20)
	if p != -1 {
		t.Errorf("Expected %d not found but returned %d", 20, p)
	}

	p = BinSearch(sortedSlice, 12)
	if p != -1 {
		t.Errorf("Expected %d not found but returned %d", 12, p)
	}
}
