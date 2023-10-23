package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	count "github.com/hitesh22rana/quik/pkg/count"
	prime "github.com/hitesh22rana/quik/pkg/prime"
	queue "github.com/hitesh22rana/quik/pkg/queue"
)

func ConcurrentQueue() {
	var wgE sync.WaitGroup
	var wgD sync.WaitGroup

	var size int = 1000000

	cq := queue.ConcurrentQueue{}

	for i := 0; i < size; i++ {
		wgE.Add(1)
		go func() {
			cq.Enqueue(rand.Int())
			wgE.Done()
		}()
		wgE.Wait()
	}

	fmt.Printf("After enqueueing %d elements, size is %d\n", size, cq.Size())

	for i := 0; i < size; i++ {
		wgD.Add(1)
		go func() {
			_, err := cq.Dequeue()
			if err != nil {
				panic(err)
			}
			wgD.Done()
		}()
		wgD.Wait()
	}

	fmt.Printf("After dequeuing %d elements, size is %d\n", size, cq.Size())
}

func ConcurrentCounter() {
	var wgE sync.WaitGroup
	var wgD sync.WaitGroup

	var counter int = 1000000

	cntr := count.Counter{}

	for i := 0; i < counter; i++ {
		wgE.Add(1)
		go func() {
			cntr.Increase()
			wgE.Done()
		}()
		wgE.Wait()
	}

	fmt.Printf("After increasing %d times, value is %d\n", counter, cntr.Value())

	for i := 0; i < counter; i++ {
		wgD.Add(1)
		go func() {
			cntr.Decrease()
			wgD.Done()
		}()
		wgD.Wait()
	}

	fmt.Printf("After decreasing %d times, value is %d\n", counter, cntr.Value())
}

func ConcurrentPrime() {
	start := time.Now()
	var maxInt int64 = 10000000
	var concurrency int8 = 12
	var wg sync.WaitGroup

	fmt.Printf("Total number of prime numbers till %d is %d\n", maxInt, prime.TotalPrimes(maxInt, concurrency, &wg))
	fmt.Printf("Finished in %s\n", time.Since(start))
}

func main() {
	// ConcurrentQueue()
	// ConcurrentCounter()

	ConcurrentPrime()
}
