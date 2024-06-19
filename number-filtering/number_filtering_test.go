package numberfiltering

import (
	"errors"
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

var isMultipleOf = func (num int, divisor int) (bool, error) {
    if divisor == 0 {
        return false, errors.New("cannot divide by zero")
    }
    return num % divisor == 0, nil
}

func isGreaterThan(threshold int) func(int) bool {
	return func(num int) bool {
		return num > threshold
	}
}

func isLessThan(threshold int) func(int) bool {
	return func(num int) bool {
		return num < threshold
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

func TestGetOddPrimeNumbers(t *testing.T) {

	t.Run("Test get odd prime numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= FilterNumbersWithAllConditions(numbers, isPrime, isOdd)
		want:= []int{3, 5, 7}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd prime numbers with [2 81 69 4 47 54 84 37 87]", func(t *testing.T) {
		numbers := []int{2, 81, 69, 4, 47, 54, 84, 37, 87}
		got:= FilterNumbersWithAllConditions(numbers, isPrime, isOdd)
		want:= []int{47, 37}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})
}

func TestGetEvenMultiplesOfFiveNumbers(t *testing.T) {
	t.Run("Test get even multiples of five numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := FilterNumbersWithAllConditions(numbers, isEven, func(num int) bool {
			multiple, _ := isMultipleOf(num, 5)
			return multiple
		})
		want := []int{10}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})

	t.Run("Test get even multiples of five numbers with [12 15 18 20 22 25 27 30]", func(t *testing.T) {
		numbers := []int{12, 15, 18, 20, 22, 25, 27, 30}
		got := FilterNumbersWithAllConditions(numbers, isEven, func(num int) bool {
			multiple, _ := isMultipleOf(num, 5)
			return multiple
		})
		want := []int{20, 30}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
}

func TestGetOddMultiplesOfThreeGreaterThanTenNumbers(t *testing.T) {

t.Run("Test get odd multiples of three and greater than ten numbers with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
		got := FilterNumbersWithAllConditions(numbers, isOdd, func(num int) bool {
			multiple, _ := isMultipleOf(num, 3)
			return multiple
		}, isGreaterThan(10))
		want := []int{15, 21}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})	

	t.Run("Test get odd multiples of three and greater than ten numbers with [12 15 18 20 22 25 27 30]", func(t *testing.T) {
		numbers := []int{12, 15, 18, 20, 22, 25, 27, 30}
		got := FilterNumbersWithAllConditions(numbers, isOdd, func(num int) bool {
			multiple, _ := isMultipleOf(num, 3)
			return multiple
		}, isGreaterThan(10))
		want := []int{15, 27}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
}

func TestStorySevenAllConditionsSample(t *testing.T){
	t.Run("Test get odd multiples of three and greater than five numbers with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		got := FilterNumbersWithAllConditions(numbers, isOdd, func(num int) bool {
			multiple, _ := isMultipleOf(num, 3)
			return multiple
		}, isGreaterThan(5))
		want := []int{9, 15}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
	
	t.Run("Test get even multiples of three and less than fifteen numbers with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		got := FilterNumbersWithAllConditions(numbers, isEven, func(num int) bool {
			multiple, _ := isMultipleOf(num, 3)
			return multiple
		}, isLessThan(15))
		want := []int{6, 12}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
	
}

func TestStoryEightAnyConditionSample(t *testing.T){
	t.Run("Test get prime or multiples of five or greater than fifteen numbers with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		got := FilterNumbersWithAnyCondition(numbers, isPrime, func(num int) bool {
			multiple, _ := isMultipleOf(num, 5)
			return multiple
		}, isGreaterThan(15))
		want := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
	
	t.Run("Test get less than 6 or multiples of three numbers with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		got := FilterNumbersWithAnyCondition(numbers, func(num int) bool {
			multiple, _ := isMultipleOf(num, 3)
			return multiple
		}, isLessThan(6))
		want := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted: %v but got: %v", want, got)
		}
	})
	
}