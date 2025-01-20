package urls

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	tree, err := html.Parse(reader)
	if err != nil {
		return nil, err
	}
	ret := []string{}
	URL, _ := url.Parse(rawBaseURL)
	ret, err = check_children_r(tree, URL)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func check_children_r(n *html.Node, URL *url.URL) ([]string, error) {
	ret := []string{}
	for v := range n.ChildNodes() {

		if v.Data == "a" {
			//fmt.Println(v.Data)
			//fmt.Println(v.Attr[0].Val)
			nu, _ := NormalizeURL(v.Attr[0].Val)
			U, _ := url.Parse(v.Attr[0].Val)
			if U.Host == "" {
				nu = URL.Host + nu
			}

			ret = append(ret, fmt.Sprintf("%v://%v", URL.Scheme, nu))
		}
		if v.FirstChild != nil {

			re, err := check_children_r(v, URL)
			if err != nil {
				return nil, err
			}
			ret = append(ret, re...)
		}

	}
	return ret, nil
}
