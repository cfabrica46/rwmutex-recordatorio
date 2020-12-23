package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()

	var (
		wg sync.WaitGroup
		m  sync.RWMutex
	)

	var a int

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			time.Sleep(time.Millisecond * 50)
			m.Lock()
			a++
			m.Unlock()

		}(&wg, &m)
	}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		time.Sleep(time.Millisecond * 50)

		go func(wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()

			fmt.Println(a)

		}(&wg, &m)
	}

	wg.Wait()

	fmt.Println("Value:", a)
	fmt.Println("Duration:", time.Since(now))

}
