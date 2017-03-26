package main

import (
	"bb/crawler"
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting BB Crawler...")
	fmt.Println("......................")

	var wg sync.WaitGroup

	wg.Add(1)
	go crawler.CrawlDressabelle(&wg)

	wg.Wait()

	fmt.Println("......................")
	fmt.Println("Finished BB Crawler!")
}
