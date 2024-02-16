package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	count := 0
	numMap := make([]bool, n)
	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= sqrt; i++ {
		if numMap[i] {
			continue
		}

		for j := i * 2; j < n; j += i {
			if !numMap[j] {
				count++
				numMap[j] = true
			}
		}

	}

	count = n - count - 2
	if count > 0 {
		return count
	}
	return 0
}

func main() {
	num := 367184
	num = 10
	count := countPrimes(num)

	fmt.Printf("\nNumber of primes less than %v are %v\n", num, count)
}
