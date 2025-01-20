package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"

	"github.com/serux/webcrawler/urls"
)

func main() {
	a := os.Args[1:]
	if len(a) < 2 {
		fmt.Println("no website and N provided ")
		os.Exit(1)
	}

	num, err := strconv.Atoi(a[1])
	if err != nil {
		fmt.Println("second param must be a number")
		os.Exit(1)
	}

	if len(a) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	ur, _ := url.Parse(a[0])
	fmt.Println("starting crawl of:", a[0])

	cfg := urls.Config{
		Pages:              map[string]int{},
		BaseURL:            ur,
		Mu:                 &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, num),
		Wg:                 &sync.WaitGroup{},
	}
	cfg.CrawlPage(a[0])
	cfg.Wg.Wait()
	for k, v := range cfg.Pages {
		fmt.Println("K: ", k, " - V: ", v)
	}

}
