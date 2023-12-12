package sorting

import (
	"sync"
	"time"
)

func Thread(el uint, wg *sync.WaitGroup, ch chan uint) {
	defer wg.Done()
	time.Sleep(time.Duration(el) * (time.Second))
	ch <- el
}

func SleepSort(arr []uint) []uint {
	var sorted []uint
	var wg sync.WaitGroup
	ch := make(chan uint)

	for _, e := range arr {
		wg.Add(1)
		go Thread(e, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for el := range ch {
		sorted = append(sorted, el)
	}

	return sorted
}
