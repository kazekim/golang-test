/*
  GoLang code created by Jirawat Harnsiriwatanakit https://github.com/kazekim
*/

package main

import (
	"fmt"
	"net/http"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("http://jirawat.kim/")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "card-body"
		}
		return false
	}
	// grab all articles and print them
	articles := scrape.FindAll(root, matcher)
	for i, article := range articles {
		fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}

	resp, err = http.Get("http://jirawat.kim/2020/01/18/golang-how-to-use-panic-and-recover-to-handle-error/")
	if err != nil {
		panic(err)
	}
	root, err = html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher = func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.P && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "post-excerpt"
		}
		return false
	}
	// grab all articles and print them
	articles = scrape.FindAll(root, matcher)
	for i, article := range articles {
		fmt.Printf("%2d %s\n", i, scrape.Text(article))
	}

}