package commands

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
	return nil
}
