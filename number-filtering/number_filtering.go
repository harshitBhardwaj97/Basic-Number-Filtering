package numberfiltering

import (
	"math"
)


func GetOddNumbers(numbers []int) []int {
	var oddNumbers []int
	for _, num := range numbers {
		if num % 2 == 1 {
			oddNumbers = append(oddNumbers, num)
		} 
	}
	return oddNumbers
}

func GetEvenNumbers(numbers []int) []int {
	var evenNumbers []int
	for _, num := range numbers {
		if num % 2 == 0 {
			evenNumbers = append(evenNumbers, num)
		} 
	}
	return evenNumbers
}

func GetPrimeNumbers(numbers []int) []int {
var primeNumbers []int
	for _, num := range numbers {
		if CheckPrime(num) {
			primeNumbers = append(primeNumbers, num)
		} 
	}
	return primeNumbers	
}

func GetOddPrimeNumbers(numbers []int) []int {
var oddPrimeNumbers []int
	for _, num := range numbers {
		if CheckPrime(num) && num % 2 == 1 {
			oddPrimeNumbers = append(oddPrimeNumbers, num)
		} 
	}
	return oddPrimeNumbers	
}

func GetEvenNumbersWhichAreMultipleOf5(numbers []int) []int {
var evenMultiplesOfFive []int
	for _, num := range numbers {
		if num % 2 == 0 && num % 5 == 0 {
			evenMultiplesOfFive = append(evenMultiplesOfFive, num)
		} 
	}
	return evenMultiplesOfFive	
}

func GetOddNumbersWhichAreMultipleOf3AndGreaterThan10(numbers []int) []int {
var oddMultiplesofThreeAndGreaterThanTen []int
	for _, num := range numbers {
		if num % 2 == 1 && num % 3 == 0 && num > 10 {
			oddMultiplesofThreeAndGreaterThanTen = append(oddMultiplesofThreeAndGreaterThanTen, num)
		} 
	}
	return oddMultiplesofThreeAndGreaterThanTen	
}

func CheckPrime(number int) bool {
	if number <= 1 {
		return false
	} else if number == 2 {
		return true
	} else  {
		isPrime:= true
		for i:=2; float64(i) <= math.Sqrt(float64(number)); i++ {
			if number % i == 0 {
				isPrime = false
				break
			}
		}
		return isPrime
	}
}