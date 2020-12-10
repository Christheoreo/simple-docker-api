package prime

import (
	"math"
	"sort"
)

var primes []int64 = []int64{}

// IsNumberPrime returns whether the given number is in fact a prime
func IsNumberPrime(number int64) bool {
	if number%3 == 0 {
		return false
	}
	sqroot := int64(math.Floor(math.Sqrt(float64(number))))

	if number%int64(sqroot) == 0 {
		return false
	}
	if len(primes) < 1 || primes[len(primes)-1] < sqroot {
		getPrimesLessThan(sqroot)
	} else {
		primes = filter(primes, func(v int64) bool {
			return int64(v) < number
		})
	}

	sort.Slice(primes, func(i, j int) bool {
		return primes[i] < primes[j]
	})
	for _, val := range primes {
		if number%val == 0 {
			return false
		}
	}
	return true
}
func getPrimesLessThan(max int64) {
	numberList := make([]int64, 0)
	var i int64 = 3
	for i = 3; i < max; i += 2 {
		if i < 9 {
			numberList = append(numberList, i)
			continue
		}
		var isPrime bool = true
		var j int64 = 3
		for j = 3; j < i-1; j += 2 {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			numberList = append(numberList, i)
		}
	}
	primes = numberList
}

// filter filters the the int slice
func filter(vs []int64, f func(int64) bool) []int64 {
	vsf := make([]int64, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
