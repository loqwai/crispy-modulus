package game

// Game represents an instance of the game
type Game interface {
	ComputeFirstPlayer()
	Draw() error
	IsDone() bool
	SetState(state State)
	Start() error
	State() State
	String() string
	Steal(card int) error
}

// State describe what state the game is currently in. It is serializable
type State struct {
	CardCount     int
	CurrentPlayer int
	Players       []PlayerState
}

type _Game struct {
	cardCount     int
	currentPlayer int
	players       []*Player
}

// New returns a new Game instance
func New(cardCount int) Game {
	return &_Game{
		cardCount:     cardCount,
		currentPlayer: 0,
		players: []*Player{
			NewPlayer(cardCount),
			NewPlayer(cardCount),
		},
	}
}

func (g *_Game) State() State {
	playerStates := make([]PlayerState, len(g.players))
	for i, p := range g.players {
		playerStates[i] = p.State()
	}
	return State{
		CardCount:     g.cardCount,
		CurrentPlayer: g.currentPlayer,
		Players:       playerStates,
	}
}

func (g *_Game) ComputeFirstPlayer() {
	currentPlayer := 0
	maxScore := 0

	for i, p := range g.players {
		if p.Score() > maxScore {
			maxScore = p.Score()
			currentPlayer = i
		}
	}
	g.currentPlayer = currentPlayer
}

func (g *_Game) Start() error {
	g.Draw()
	g.Draw()
	return nil
}

func (g *_Game) SetState(state State) {
	g.cardCount = state.CardCount
	g.currentPlayer = state.CurrentPlayer
	g.players = make([]*Player, len(state.Players))
	for i, p := range state.Players {
		g.players[i] = NewPlayerFromState(state.CardCount, p)
	}
}

func (g *_Game) String() string {
	return `
          | Dealt | Remaining | Total | Modulus |
  Player1 | 2 3 4 |       1 5 |     9 |       4 |
  Player2 | 1 4 5 |       2 3 |    10 |       0 |
	`
}

func (g *_Game) Draw() error {
	player := g.players[g.currentPlayer]
	err := player.Draw()
	if err != nil {
		return err
	}

	g.currentPlayer = (g.currentPlayer + 1) % len(g.players)
	return nil
}

func (g *_Game) IsDone() bool {
	if len(g.State().Players[0].Deck) == 0 {
		return true
	}
	if len(g.State().Players[1].Deck) == 0 {
		return true
	}
	return false
}

func (g *_Game) Steal(card int) error {
	otherPlayerIndex := (g.currentPlayer + 1) % len(g.players)

	err := g.players[otherPlayerIndex].Steal(card)
	if err != nil {
		return err
	}

	g.players[g.currentPlayer].Give(card)

	g.currentPlayer = (g.currentPlayer + 1) % len(g.players)
	return nil
}
