package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {

	var target string = "https://www.basketball-reference.com/boxscores/202101120BRK.html"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("div.table_container", func(e *colly.HTMLElement) {
		if strings.TrimSpace(e.ChildText("table caption")) == "Table" {
			e.ForEachWithBreak("table tbody", printStatHeader)
		}
	})

	c.Visit(target)
}

func printStatHeader(i int, t *colly.HTMLElement) bool {
	t.ForEach("tr td", func(j int, td *colly.HTMLElement) {
		fmt.Print(td.Attr("data-stat"), " / ")
	})
	return false
}
