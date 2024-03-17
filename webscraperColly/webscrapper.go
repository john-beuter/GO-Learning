package main

import (
	"fmt"

	"github.com/gocolly/colly" //I downloaded you though?
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	//Scraping HTML
	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttr("a", "href")
		fmt.Println(links)
	})

	c.Visit("https://en.wikipedia.org/wiki/Dune_(franchise)")
}
