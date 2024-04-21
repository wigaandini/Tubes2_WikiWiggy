package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func validLink(link string) bool {
	invalidPrefixes := []string{"/wiki/Special:", "/wiki/Talk:", "/wiki/User:", "/wiki/Portal:", "/wiki/Wikipedia:", "/wiki/File:", "/wiki/Category:", "/wiki/Help:"}
	for _, prefix := range invalidPrefixes {
		if strings.HasPrefix(link, prefix) {
			return false
		}
	}
	return strings.HasPrefix(link, "/wiki/")
}

func main() {
	link := "https://en.wikipedia.org/wiki/Artificial_intelligence"
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	linkCount := 0

	c.OnHTML("div#mw-content-text "+"a[href]", func(e *colly.HTMLElement) {
		neighborLink := e.Attr("href")
		if validLink(neighborLink) {
			linkCount++
		}
	})

	c.Visit(link)

	fmt.Println("Number of links:", linkCount)
}
