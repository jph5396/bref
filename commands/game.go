package commands

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/jph5396/bref/model"
)

//BaseURL base url to strape.
var BaseURL string = "https://www.basketball-reference.com"

//GetGame returns the box score from the game.
func GetGame(gameid string) model.Game {
	var g model.Game
	g.ID = gameid

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("div.scorebox", func(e *colly.HTMLElement) {
		g.GameInfo = parseScorebox(e)

	})

	c.OnHTML("div.table_container", func(e *colly.HTMLElement) {

		//if this element represents the away basic table.
		if e.Attr("id") == "div_box-"+g.GameInfo.Away.Initals+"-game-basic" {
			g.AwayBox.Basic = parseBaseTable(e)
		}
		//if this element represents the away advanced table.
		if e.Attr("id") == "div_box-"+g.GameInfo.Away.Initals+"-game-advanced" {
			g.AwayBox.Advanced = parseAdvTable(e)
		}

		// home version.
		if e.Attr("id") == "div_box-"+g.GameInfo.Home.Initals+"-game-basic" {
			g.HomeBox.Basic = parseBaseTable(e)
		}

		if e.Attr("id") == "div_box-"+g.GameInfo.Home.Initals+"-game-advanced" {
			g.HomeBox.Advanced = parseAdvTable(e)
		}
	})

	c.Visit(BaseURL + "/boxscores/" + gameid + ".html")

	return g
}
