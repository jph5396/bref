package commands

type (
	//GameCommand command that is capable of scraping an entire day.
	GameCommand struct {
		Name   string
		Config CommandConfig
		Args   []string
	}
)

//NewGameCommand returns a GameCommand type implementing the Commmand interface.
func NewGameCommand() GameCommand {
	return GameCommand{
		Name: "game",
	}
}

//CommandName returns name
func (gc *GameCommand) CommandName() string {
	return gc.Name
}

// CommandDescription describe command.
func (gc *GameCommand) CommandDescription() string {
	return `
	game <gameid > - scrape the boxscore of a specific game. the gameid can be found in the last 
	section of the boxscore URL before the -.html suffix.
	`
}

//SetConfig for the command
func (gc *GameCommand) SetConfig(c CommandConfig) {
	gc.Config = c
}

//Run Execute the command
func (gc *GameCommand) Run(args []string) error {

	g, err := GetGame(args[0])
	if err != nil {
		return err
	}
	filename := g.ID + ".json"
	if gc.Config.SaveDir != "" {
		filename = gc.Config.SaveDir + filename
	}

	err = JSONFileWriter(filename, g)
	if err != nil {
		return err
	}
	return nil
}
