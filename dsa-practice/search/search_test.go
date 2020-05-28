package search

import "testing"

func TestLinearSearch(t *testing.T) {
	input := []int{3, 5, 7, 11, 15, 27}
	result, _ := linearsearch(input, 15)
	expected := 4
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
}

func TestBinarySearch(t *testing.T) {
	input := []int{3, 5, 7, 11, 15, 27}
	result, err := binarysearch(input, 3, 0, 5)
	expected := 0
	if result != expected && err != nil {
		t.Errorf("expected %v got %v", expected, result)
	}
	result, _ = binarysearch(input, 15, 0, 5)
	expected = 4
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
	_, err = binarysearch(input, -1, 0, 5)
	if err == nil {
		t.Errorf("error should have been thrown for element not in slice")
	}
}

func TestJumpSearch(t *testing.T) {
	input := []int{3, 5, 7, 11, 15, 27}
	result, err := jumpsearch(input, 3, 2)
	expected := 0
	if result != expected && err != nil {
		t.Errorf("expected %v got %v", expected, result)
	}
	result, _ = jumpsearch(input, 15, 1)
	expected = 4
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
	result, _ = jumpsearch(input, 15, 100)
	expected = 4
	if result != expected {
		t.Errorf("expected %v got %v", expected, result)
	}
	_, err = jumpsearch(input, 15, -1)
	if err == nil {
		t.Errorf("error should have been thrown for element not in slice")
	}
	_, err = jumpsearch(input, -1, 3)
	if err == nil {
		t.Errorf("error should have been thrown for element not in slice")
	}
}
