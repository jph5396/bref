package commands

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/jph5396/bref/model"
)

//this file contains common functions that are used to parse data from colly html objects and
// return the data in a way that matches our models. this way the can be reused accross commands that
// might have similar pieces returned.

//ParseScorebox should be used with div.scorebox as the selector
func ParseScorebox(e *colly.HTMLElement) model.ScoreBox {
	var scorebox model.ScoreBox

	//get team names
	names := e.ChildTexts("div strong a")
	scorebox.Away.Name = names[0]
	scorebox.Home.Name = names[1]
	//get scores
	scores := e.ChildTexts("div.score")
	scorebox.Away.Score, _ = strconv.Atoi(scores[0])
	scorebox.Home.Score, _ = strconv.Atoi(scores[1])

	meta := e.ChildTexts("div.scorebox_meta div")
	scorebox.Time = meta[0]
	scorebox.Location = meta[1]

	return scorebox
}

func parseAdvTable(e *colly.HTMLElement) []model.PlayerAdvBox {
	var advancedArray []model.PlayerAdvBox
	e.ForEach("table tbody tr", func(i int, ele *colly.HTMLElement) {
		var player model.PlayerAdvBox
		player.Name = ele.ChildText("th a")
		player.ScoureID = ele.ChildAttr("th", "data-append-csv")
		ele.ForEach("td", func(j int, td *colly.HTMLElement) {
			//skip blank fields.
			if td.Text != "" {
				err := player.AddByTag(td.Attr("data-stat"), td.Text)
				if err != nil {
					fmt.Printf("data-row: %v, stat: %v, Error: %v", i, j, err.Error())
				}
			}
		})
		if !reflect.ValueOf(player).IsZero() {
			advancedArray = append(advancedArray, player)
		}
	})

	//repeat the same process for team totals stored in table footer
	var teamTotal model.PlayerAdvBox
	e.ForEach("table tbody tfoot tr", func(i int, ele *colly.HTMLElement) {
		teamTotal.Name = ele.ChildText("th.left")
		ele.ForEach("td", func(j int, td *colly.HTMLElement) {
			//skip blank fields.
			if td.Text != "" {
				err := teamTotal.AddByTag(td.Attr("data-stat"), td.Text)
				if err != nil {
					fmt.Printf("data-row: %v, stat: %v, Error: %v", i, j, err.Error())
				}
			}
		})
		teamTotal.ScoureID = "TEAM"
		advancedArray = append(advancedArray, teamTotal)
	})
	return advancedArray
}
