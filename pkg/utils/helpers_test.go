package utils

import "testing"

func TestEach(t *testing.T) {
	// Test case 1: Empty list
	var list []int
	var fn func(int)
	Each(list, fn)

	// Test case 2: Non-empty list
	list = []int{1, 2, 3, 4, 5}
	fn = func(value int) {
		if value < 0 || value > 5 {
			t.Errorf("Invalid value: %d", value)
		}
	}
	Each(list, fn)

	// // Test case 3: Nil list
	list = nil
	fn = func(value int) {
		t.Errorf("Invalid value: %d", value)
	}
	Each(list, fn)

	// Test case 4: Nil function
	list = []int{1, 2, 3, 4, 5}
	fn = nil
	Each(list, fn)
}
