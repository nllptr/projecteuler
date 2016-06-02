package projecteuler

// DoSum sums up all multiples of 3 or 5 below limit
func DoSum(limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
