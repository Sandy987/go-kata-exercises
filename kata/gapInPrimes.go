package kata

import "math"

// Gap returns the first pair of two prime numbers spaced with a gap of g between the limits m, n if these numbers exist
func Gap(g, m, n int) []int {
	currentPrime := -1
	for i := m; i <= n; i++ {
		if isPrime(i) {
			if currentPrime > 0 && i-currentPrime == g {
				return []int{currentPrime, i}
			}
			currentPrime = i
		}
	}
	return nil
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
