package hellomath

// Sums all elements in an array of integers
func SumArray(array []int) int {
	sum := 0

	for _, v := range array {
		sum += v
	}
	return sum
}
