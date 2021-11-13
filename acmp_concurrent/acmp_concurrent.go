package acmp_concurrent

import (
	acmp "../acmp"
	"fmt"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	var sMap sync.Map
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string){
			defer wg.Done()
			diff := acmp.Difficulty(url)
			sMap.Store(url, diff)
		}(url)
	}
	wg.Wait()

	result := make(map[string]float64)
	for _, url := range urls {
		diff, ok := sMap.Load(url)
		if !ok {
			fmt.Println(diff, ok)
			break
		}
		result[url] = diff.(float64)
	}
	return result
}


