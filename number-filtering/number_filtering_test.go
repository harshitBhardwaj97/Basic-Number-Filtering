package numberfiltering

import (
	"reflect"
	"testing"
)

// Define more conditions here
var isOdd = func (num int) bool {
	return num % 2 == 1
}

func TestGetOddNumbers(t *testing.T) {

	t.Run("Test get odd numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= FilterNumbersWithAllConditions(numbers, isOdd)
		want:= []int{1, 3, 5, 7, 9}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd numbers with [12 77 99 60 28 93 75 49]", func(t *testing.T) {
		numbers := []int{12, 77, 99, 60, 18, 93, 75, 49}
		got:= FilterNumbersWithAllConditions(numbers, isOdd)
		want:= []int{77, 99, 93, 75, 49}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})
}