package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/jph5396/bref/model"
)

//BaseURL base url to strape.
var BaseURL string = "https://www.basketball-reference.com"

//GetGame returns the box score from the game.
func GetGame(gameid string) (model.Game, error) {
	var g model.Game
	var returnErr error = nil
	g.ID = gameid

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		returnErr = fmt.Errorf(" failed when getting %v, status: %v, error: %v", gameid, r.StatusCode, err.Error())
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

	return g, returnErr
}

// JSONFileWriter reusable function for writing structs to json files.
// will return any errors that occur, or nil if it succeeds.
func JSONFileWriter(path string, data interface{}) error {

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	encodeErr := encoder.Encode(data)
	if encodeErr != nil {
		return encodeErr
	}

	defer file.Close()
	fmt.Println("Created file: ", path)
	return nil
}
