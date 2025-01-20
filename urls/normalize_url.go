package urls

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type Config struct {
	Pages              map[string]int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
}

func NormalizeURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host + u.Path, nil
}

func GetHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("server Error")
	}
	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("not text/html")
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	if !strings.Contains(rawCurrentURL, cfg.BaseURL.String()) {
		return
	}

	isFirst := cfg.addPageVisit(rawCurrentURL)
	if !isFirst {
		return
	}

	fmt.Println("Crawling: ", rawCurrentURL)
	webHtml, err := GetHTML(rawCurrentURL)
	if err != nil {
		fmt.Println(err)

	}

	weblinks, err := GetURLsFromHTML(webHtml, rawCurrentURL)
	if err != nil {
		return
	}

	for _, v := range weblinks {
		cfg.Wg.Add(1)
		go func() {
			defer func() {
				cfg.Wg.Done()
				<-cfg.ConcurrencyControl
			}()
			cfg.ConcurrencyControl <- struct{}{}
			cfg.CrawlPage(v)
		}()
	}

}

func (cfg *Config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()

	_, ok := cfg.Pages[normalizedURL]
	if ok {
		cfg.Pages[normalizedURL]++
		return false
	}
	cfg.Pages[normalizedURL] = 1
	return true

}
