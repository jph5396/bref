package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jph5396/bref/model"

	"github.com/gocolly/colly/v2"
)

type (
	//DayCommand command that is capable of scraping an entire day.
	DayCommand struct {
		Name   string
		Config CommandConfig
		Args   []string
	}
)

//NewDayCommand returns a new day command implementing the Command interface.
func NewDayCommand() DayCommand {
	return DayCommand{
		Name: "day",
	}
}

//CommandName returns name
func (dc *DayCommand) CommandName() string {
	return dc.Name
}

// CommandDescription describe command.
func (dc *DayCommand) CommandDescription() string {
	return `
	day <date: YYYY-MM-DD> - the day command will scrape all boxscores that occured for that day`
}

//SetConfig for the command
func (dc *DayCommand) SetConfig(c CommandConfig) {
	dc.Config = c
}

//Run Execute the command
func (dc *DayCommand) Run(args []string) error {
	rawdate := args[0]
	pattern, err := regexp.Compile(`^\d{4}-\d{2}-\d{2}$`)
	if err != nil {
		return err
	}

	if !pattern.MatchString(rawdate) {
		return fmt.Errorf("%v does not match YYYY-MM-DD format", rawdate)
	}

	date := strings.Split(rawdate, "-")
	var gameids []string

	c := colly.NewCollector()
	var returnErr error = nil
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		returnErr = fmt.Errorf(" failed status: %v, error: %v", r.StatusCode, err.Error())
	})

	c.OnHTML("div.game_summary p.links a", func(e *colly.HTMLElement) {
		if e.Text == "Box Score" {
			href := strings.Split(e.Attr("href"), "/")
			gameid := strings.Split(href[len(href)-1], ".")[0]
			gameids = append(gameids, gameid)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Printf("found %v games on %v \n", len(gameids), rawdate)
		for _, val := range gameids {
			g, err := GetGame(val)
			if err != nil {
				returnErr = err
			}

			dc.Save(g)
		}
	})

	c.Visit(BaseURL + fmt.Sprintf("/boxscores?month=%v&day=%v&year=%v", date[1], date[2], date[0]))

	return returnErr
}

//Save function to be called to save a game.
func (dc *DayCommand) Save(g model.Game) error {

	filename := g.ID + ".json"
	if dc.Config.SaveDir != "" {
		filename = dc.Config.SaveDir + filename
	}

	err := JSONFileWriter(filename, g)
	if err != nil {
		return err
	}
	return nil
}
