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
		Name    string
		Score   int
		Initals string
	}

	//BoxScore contains both the basic and advanced stats for each player on a team.
	BoxScore struct {
		Basic    []PlayerBasicBox
		Advanced []PlayerAdvBox
	}

	//Game represents the entire game.
	Game struct {
		ID       string
		GameInfo ScoreBox
		AwayBox  BoxScore
		HomeBox  BoxScore
	}
)
