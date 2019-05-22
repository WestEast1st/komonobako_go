package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

type Result struct {
	Tag  string
	Attr []html.Attribute
}

func main() {
	f, _ := os.Open("test.html")
	r := bufio.NewReader(f)
	/*
		ParseItem(r)
	*/
}

func ParseItem(r io.Reader) []Result {
	results := []Result{}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println(err)
	}
	var f func(*html.Node)
	elementfillter := func(d string) bool {
		for _, v := range []string{"a", "form", "img", "input"} {
			if d == v {
				return true
			}
		}
		return false
	}

	resstr := func(d *html.Node) {
		results = append(results, Result{
			Tag:  d.Data,
			Attr: d.Attr,
		})
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && elementfillter(n.Data) {
			resstr(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println(results)
	return results
}
