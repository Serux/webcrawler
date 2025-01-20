package urls

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NormalizeURL(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return u.Host + u.Path, nil
}

func GetHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("server Error")
	}
	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("not text/html")
	}
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil

}
