package iterator

import (
	"testing"
)

func TestStringsIterator(t *testing.T) {

	stringSlice := []string{"mango", "apple", "grapes", "banana", "pears"}

	var expectedIndx = 4

	resultIndx := IndexOfString(stringSlice, "pears")

	if resultIndx != expectedIndx {
		t.Errorf("Expected %d, but got %d", expectedIndx, resultIndx)
	}

	expectedIndx = 2

	resultIndx = IndexOfString(stringSlice, "grapes")

	if resultIndx != expectedIndx {
		t.Errorf("Expected %d, but got %d", expectedIndx, resultIndx)
	}

	exists := IsStringIncluded(stringSlice, "banana")

	if !exists {
		t.Errorf("Expected true, but got false")
	}

	match := MatchAnyString(stringSlice, func(s string) bool {
		return len(s) == 6
	})

	if !match {
		t.Errorf("Expected true, but got false")
	}

	match = MatchAllStrings(stringSlice, func(s string) bool {
		return len(s) <= 6
	})

	if !match {
		t.Errorf("Expected true, but got false")
	}

	expectedSlice := []string{"mango", "apple", "pears"}

	resultedSlice := FilterStrings(stringSlice, func(s string) bool {
		return len(s) == 5
	})

	if len(expectedSlice) != len(resultedSlice) {
		t.Errorf("Filtered strings were not expected")
	}

	for _, v := range resultedSlice {
		if v != "mango" && v != "apple" && v != "pears" {
			t.Errorf("Filtered string was not expected")
		}
	}
}
