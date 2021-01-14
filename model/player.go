package model

type (
	//PlayerBasicBox represents a players basic statlines in a game
	PlayerBasicBox struct {
		Name      string
		MP        string
		FG        int
		FGA       int
		FGPct     float64
		ThreeP    int
		ThreePA   int
		ThreePPct float64
		FT        int
		FTA       int
		FTPct     float64
		ORB       int
		DRB       int
		TRB       int
		AST       int
		STL       int
		BLK       int
		TOV       int
		PF        int
		PTS       int
		PlusMin   string
	}

	//PlayerAdvBox represents a players advanced stats during a game.
	PlayerAdvBox struct {
		Name         string
		MP           string
		TrueShootPct float64
		ThreePARate  float64
		FTARate      float64
		ORBPct       float64
		DRBPct       float64
		TRBPct       float64
		ASTPct       float64
		STLPct       float64
		BLKPct       float64
		TOVPct       float64
		USGRate      float64
		ORTG         int
		DRTG         int
		BPM          string
	}
)
