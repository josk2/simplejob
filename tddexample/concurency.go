package tddexample

import (
	"sync"
)

type WebsiteChecker func(string) bool

func CheckWebsite(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool)

	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(u string, wg *sync.WaitGroup) {
			result[u] = wc(u)
			wg.Done()
		}(url, &wg)
	}

	wg.Wait()

	return result
}
