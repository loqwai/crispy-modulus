package game

//Player represents the data of a single player.
type Player struct {
	MyCards []int8
}

// Game represents an instance of the game
type Game struct {
	//the gamestate of player 1
	Players []Player
}

// New returns a new Game instance
func New() Game {
	return Game{
		Players: []Player{
			NewPlayer(),
			NewPlayer(),
		},
	}
}

//NewPlayer returns a new Player instance
func NewPlayer() Player {
	return Player{
		MyCards: []int8{
			0,
			0,
			0,
			0,
			0,
		},
	}
}

func (g *Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}
