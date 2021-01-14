package commands

import (
	"strconv"

	"github.com/gocolly/colly/v2"
	"github.com/jph5396/bref/model"
)

//this file contains common functions that are used to parse data from colly html objects and
// return the data in a way that matches our models. this way the can be reused accross commands that
// might have similar pieces returned.

//parseScorebox should be used with div.scorebox as the selector
func parseScorebox(e *colly.HTMLElement) model.ScoreBox {
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
