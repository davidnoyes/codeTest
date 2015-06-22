package codeTest

func MultiplesOf3And5(maxValue int) []int {
	results := []int{}

	for i := 1; i < maxValue; i++ {
		if i%3 == 0 || i%5 == 0 {
			results = append(results, i)
		}
	}

	return results
}

func SumMultiples(values []int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}
