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
	return "Game state!"
}
