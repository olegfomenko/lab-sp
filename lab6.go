package main

import (
	"crypto/rand"
	"sync"
)

func RunLab6() {
	const THREADS = 10
	const VALUES = 100

	wg := &sync.WaitGroup{}

	for i := 0; i < THREADS; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < VALUES; j++ {
				_, err := rand.Prime(rand.Reader, 256)
				if err != nil {
					panic(err)
				}
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
}
