package numberfiltering

import (
	"math"
	"reflect"
	"testing"
)

// Define more conditions here
var isOdd = func (num int) bool {
	return num % 2 == 1
}

var isEven = func (num int) bool {
    return num % 2 == 0
}

var isPrime = func(number int) bool {
    if number <= 1 {
        return false
    } else if number == 2 {
        return true
    } else  {
        isPrimeNumber:= true
        for i:=2; float64(i) <= math.Sqrt(float64(number)); i++ {
            if number % i == 0 {
                isPrimeNumber = false
                break
            }
        }
        return isPrimeNumber
    }
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

func TestGetEvenNumbers(t *testing.T) {

	t.Run("Test get even numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= FilterNumbersWithAllConditions(numbers, isEven)
		want:= []int{2, 4, 6, 8, 10}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get even numbers with [12 77 99 60 28 93 75 49]", func(t *testing.T) {
		numbers := []int{12, 77, 99, 60, 18, 93, 75, 49}
		got:= FilterNumbersWithAllConditions(numbers, isEven)
		want:= []int{12, 60, 18}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})
}

func TestGetPrimeNumbers(t *testing.T) {

	t.Run("Test get prime numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= FilterNumbersWithAllConditions(numbers, isPrime)
		want:= []int{2, 3, 5, 7}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get prime numbers with [20 32 59 41 60 27 50 29 61 90]", func(t *testing.T) {
		numbers := []int{20, 32, 59, 41, 60, 27, 50, 29, 61, 90}
		got:= FilterNumbersWithAllConditions(numbers, isPrime)
		want:= []int{59, 41, 29, 61}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})
}