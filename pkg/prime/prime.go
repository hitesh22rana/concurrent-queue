package prime

import (
	"math"
	"sync"
	"sync/atomic"
)

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := int64(3); i <= int64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func calculatePrimes(maxInt int64, concurrency int8, currentNum *int64, primeCount *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		num := atomic.AddInt64(currentNum, 1)

		if num > int64(maxInt) {
			break
		}

		if isPrime(num) {
			atomic.AddInt64(primeCount, 1)
		}
	}
}

func TotalPrimes(maxInt int64, concurrency int8, wg *sync.WaitGroup) int64 {
	var totalPrimeNumbers int64 = 0
	var currentNum int64 = 1

	for i := 0; i < int(concurrency); i++ {
		wg.Add(1)
		go calculatePrimes(maxInt, concurrency, &currentNum, &totalPrimeNumbers, wg)
	}

	wg.Wait()

	return totalPrimeNumbers
}
