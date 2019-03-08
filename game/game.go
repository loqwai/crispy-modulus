package game

// Game represents an instance of the card game
type Game interface {
	String() string
}

type _Game struct{}

// New returns a new Game instance
func New() Game {
	return &_Game{}
}

func (g *_Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}
