package main

import (
	"fmt"
	"strings"

	"github.com/jph5396/bref/commands"

	"github.com/gocolly/colly/v2"
)

func main() {
	// var advArr []model.PlayerAdvBox
	var target string = "https://www.basketball-reference.com/boxscores/202101130CHO.html"
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("div.scorebox", func(e *colly.HTMLElement) {
		box := commands.ParseScorebox(e)
		fmt.Println(box)
	})

	c.OnHTML("div.table_container", func(e *colly.HTMLElement) {

		if strings.TrimSpace(e.ChildText("table caption")) == "Table" {

		}
	})

	c.Visit(target)

	// for _, val := range advArr {
	// 	fmt.Println(val)
	// }
}

func printStatHeader(i int, t *colly.HTMLElement) bool {
	t.ForEach("tr td", func(j int, td *colly.HTMLElement) {
		fmt.Print(td.Attr("data-stat"), " / ")
	})
	return false
}
