package projecteuler

// IsPrime checks if a number is prime
func IsPrime(n int) bool {
	if n == 1 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// FindLargestPrime returns the largest prime number lower than limit
func FindLargestPrime(n int) int {
	remainder := n
	primes := []int{}

	for {
		for i := 2; i < remainder; i++ {
			if IsPrime(i) && remainder%i == 0 {
				primes = append(primes, i)
				remainder = remainder / i
				break
			}
		}
		if IsPrime(remainder) {
			primes = append(primes, remainder)
			break
		}
	}
	largest := 0
	for _, v := range primes {
		if v > largest {
			largest = v
		}
	}
	return largest
}
