package iterator

import (
	"testing"
)

func TestIntegerIterator(t *testing.T) {

	intSlice := []int{0, 1, 2, 99, 151}

	var expectedIndx = 4

	resultIndx := IndexOfInteger(intSlice, 151)

	if resultIndx != expectedIndx {
		t.Errorf("Expected %d, but got %d", expectedIndx, resultIndx)
	}

	expectedIndx = 2

	resultIndx = IndexOfInteger(intSlice, 2)

	if resultIndx != expectedIndx {
		t.Errorf("Expected %d, but got %d", expectedIndx, resultIndx)
	}

	exists := IsIntegerIncluded(intSlice, 151)

	if !exists {
		t.Errorf("Expected true, but got false")
	}

	match := MatchAnyInteger(intSlice, func(i int) bool {
		return i > 150
	})

	if !match {
		t.Errorf("Expected true, but got false")
	}

	match = MatchAllIntegers(intSlice, func(i int) bool {
		return i < 200
	})

	if !match {
		t.Errorf("Expected true, but got false")
	}

	intSlice1 := []int{0, 1, 2, 4, 6, 99, 88, 151}

	expectedSlice := []int{0, 2, 4, 6, 88}

	resultedSlice := FilterIntegers(intSlice1, func(i int) bool {
		return (i == 0) || (i%2 == 0)
	})

	if len(expectedSlice) != len(resultedSlice) {
		t.Errorf("Filtered integers were not expected")
	}

	for _, v := range resultedSlice {
		if v != 0 && v != 2 && v != 4 && v != 6 && v != 88 {
			t.Errorf("Filtered integer was not expected")
		}
	}

}
