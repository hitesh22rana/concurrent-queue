package main

import (
	"fmt"
	"math/rand"
	"sync"

	queue "github.com/hitesh22rana/quik/pkg"
)

var wgE sync.WaitGroup
var wgD sync.WaitGroup

func main() {
	cq := queue.ConcurrentQueue{}

	for i := 0; i < 1000000; i++ {
		wgE.Add(1)
		go func() {
			cq.Enqueue(rand.Int31())
			wgE.Done()
		}()
		wgE.Wait()
	}

	fmt.Println(cq.Size())

	for i := 0; i < 1000000; i++ {
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

	fmt.Println(cq.Size())
}
