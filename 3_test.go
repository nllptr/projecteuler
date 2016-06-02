package projecteuler

import (
	"fmt"
	"testing"
)

var isPrimeTests = []struct {
	input    int
	expected bool
}{
	{2, true},
	{4, false},
	{9, false},
	{13, true},
}

func TestIsPrime(t *testing.T) {
	for _, td := range isPrimeTests {
		actual := IsPrime(td.input)
		if actual != td.expected {
			t.Errorf("IsPrime(%d): exptected %t, actual %t", td.input, td.expected, actual)
		}
	}
}

var findLargestPrimeTests = []struct {
	input    int
	expected int
}{
	{13195, 29},
}

func TestFindLargestPrime(t *testing.T) {
	for _, td := range findLargestPrimeTests {
		actual := FindLargestPrime(td.input)
		if actual != td.expected {
			t.Errorf("FindLargestPrime(%d): expected %d, actual %d", td.input, td.expected, actual)
		}
	}
	fmt.Printf("3: FindLargestPrime(600851475143) = %d\n", FindLargestPrime(600851475143))
}
