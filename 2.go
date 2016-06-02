package projecteuler

func SumFib(limit int) int {
	sum := 0
	low := 1
	high := 2
	next := 3
	for high < limit {
		if high%2 == 0 {
			sum += high
		}
		low = high
		high = next
		next = low + high
	}
	return sum
}
