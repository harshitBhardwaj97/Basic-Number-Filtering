package numberfiltering


func FilterNumbersWithAllConditions(numbers []int, conditions ...func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		matchesAllConditions := true

		for _, condition := range conditions {
			if !condition(num) {
				matchesAllConditions = false
				break
			}
		}
		if matchesAllConditions {
            result = append(result, num)
        }
	}
	return result
}

func FilterNumbersWithAnyCondition(numbers []int, conditions ...func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		matchesAnyCondition := false

		for _, condition := range conditions {
			if condition(num) {
				matchesAnyCondition = true
			}
		}
		if matchesAnyCondition {
            result = append(result, num)
        }
	}
	return result
}