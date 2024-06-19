package numberfiltering

import (
	"reflect"
	"testing"
)

func TestGetOddNumbers(t *testing.T) {
	
	t.Run("Test get odd numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= GetOddNumbers(numbers)
		want:= []int{1, 3, 5, 7, 9}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd numbers with [12 77 99 60 28 93 75 49]", func(t *testing.T) {
		numbers := []int{12, 77, 99, 60, 18, 93, 75, 49}
		got:= GetOddNumbers(numbers)
		want:= []int{77, 99, 93, 75, 49}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestGetEvenNumbers(t *testing.T) {
	
	t.Run("Test get even numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= GetEvenNumbers(numbers)
		want:= []int{2, 4, 6, 8, 10}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get even numbers with [12 77 99 60 28 93 75 49]", func(t *testing.T) {
		numbers := []int{12, 77, 99, 60, 18, 93, 75, 49}
		got:= GetEvenNumbers(numbers)
		want:= []int{12, 60, 18}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestGetPrimeNumbers(t *testing.T) {
	
	t.Run("Test get prime numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= GetPrimeNumbers(numbers)
		want:= []int{2, 3, 5, 7}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd numbers with [81 69 4 47 54 84 37 87]", func(t *testing.T) {
		numbers := []int{81, 69, 4, 47, 54, 84, 37, 87, -3}
		got:= GetPrimeNumbers(numbers)
		want:= []int{47, 37}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestGetOddPrimeNumbers(t *testing.T) {
	
	t.Run("Test get odd prime numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= GetOddPrimeNumbers(numbers)
		want:= []int{3, 5, 7}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd numbers with [20 32 59 41 60 27 50 29 61 90]", func(t *testing.T) {
		numbers := []int{20, 32, 59, 41, 60, 27, 50, 29, 61, 90}
		got:= GetPrimeNumbers(numbers)
		want:= []int{59, 41, 29, 61}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestGetEvenFiveMultipleNumbers(t *testing.T) {
	
	t.Run("Test get even five multiple numbers with [1 2 3 4 5 6 7 8 9 10]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got:= GetEvenNumbersWhichAreMultipleOf5(numbers)
		want:= []int{10}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get even five multiple numbers with [42 24 1 93 57 3 11 80 6 59]", func(t *testing.T) {
		numbers := []int{42, 24, 1, 93, 57, 3, 11, 80, 6, 59}
		got:= GetEvenNumbersWhichAreMultipleOf5(numbers)
		want:= []int{80}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestGetOddThreeMultipleNumbersGreaterThanTen(t *testing.T) {
	
	t.Run("Test get odd three multiple numbers > 10 with [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
		got:= GetOddNumbersWhichAreMultipleOf3AndGreaterThan10(numbers)
		want:= []int{15}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

	t.Run("Test get odd three multiple numbers > 10 with [42 24 1 93 57 3 11 80 6 59]", func(t *testing.T) {
		numbers := []int{42, 24, 1, 93, 57, 3, 11, 80, 6, 59}
		got:= GetOddNumbersWhichAreMultipleOf3AndGreaterThan10(numbers)
		want:= []int{93, 57}

		if !reflect.DeepEqual(got, want) {
		t.Errorf("wanted: %v but got: %v", want, got)
	}
	})

}

func TestPrimeNumberFunction(t *testing.T) {
	t.Run("2 should be returned as prime", func (t *testing.T)  {
		num:= 2
		want:= true
		got:= CheckPrime(num)

		if got!= want {
		t.Errorf("wanted: %v but got: %v", want, got)	
		}
	})

	t.Run("9 should not be returned as prime", func (t *testing.T)  {
		num:= 9
		want:= false
		got:= CheckPrime(num)

		if got!= want {
		t.Errorf("wanted: %v but got: %v", want, got)	
		}
	})

	t.Run("104729 should be returned as prime", func (t *testing.T)  {
		num:= 104729
		want:= true
		got:= CheckPrime(num)

		if got!= want {
		t.Errorf("wanted: %v but got: %v", want, got)	
		}
	})

	
}