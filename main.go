package main

import (
	"fmt"
	"os"

	"github.com/serux/webcrawler/urls"
)

func main() {
	a := os.Args[1:]
	if len(a) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(a) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Println("starting crawl of:", a[0])
	fmt.Println(urls.GetHTML(a[0]))

}
