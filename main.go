package main

import (
	"fmt"
	"net/url"
	"os"
	"sort"
	"strconv"
	"sync"

	"github.com/serux/webcrawler/urls"
)

func main() {
	a := os.Args[1:]
	if len(a) < 3 {
		fmt.Println("no website and N provided ")
		os.Exit(1)
	}

	threads, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Println("second param must be a number")
		os.Exit(1)
	}

	max, err := strconv.Atoi(a[2])
	if err != nil {
		fmt.Println("third param must be a number")
		os.Exit(1)
	}

	if len(a) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseUrl, _ := url.Parse(a[0])

	fmt.Println("starting crawl of:", a[0])

	cfg := urls.Config{
		Pages:              map[string]int{},
		BaseURL:            baseUrl,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, threads),
		Wg:                 &sync.WaitGroup{},
		Maxpages:           max,
	}

	cfg.CrawlPage(a[0])
	cfg.Wg.Wait()

	printReport(cfg.Pages, a[0])

}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Println("=============================")
	keys := make([]string, 0, len(pages))
	for key := range pages {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return pages[keys[i]] > pages[keys[j]]
	})
	for _, v := range keys {
		fmt.Printf("Found %v internal links to %v\n", pages[v], v)
		//fmt.Println("K: ", pages[v], " - V: ", v)
	}

}
