package model

type (

	//ScoreBox contains basic data about the game.
	ScoreBox struct {
		Away     Team
		Home     Team
		Location string
		Time     string
	}
	//Team represents a team's result after a game.
	Team struct {
		Name   string
		Score  int
		Record string
	}
)
