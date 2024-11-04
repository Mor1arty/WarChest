package game

type Team int

const (
	TeamWhite Team = iota
	TeamBlack
)

type Player struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Team        Team    `json:"team"`
	Supply      []*Unit `json:"supply"`
	Hand        []*Unit `json:"hand"`
	Bag         []*Unit `json:"bag"`
	DiscardPile []*Unit `json:"discardPile"`
	Eliminated  []*Unit `json:"eliminated"`
}

func NewPlayer(id, name string, team Team, supply []*Unit, hand []*Unit, bag []*Unit, discardPile []*Unit, eliminated []*Unit) *Player {
	return &Player{
		ID:          id,
		Name:        name,
		Team:        team,
		Supply:      supply,
		Hand:        hand,
		Bag:         bag,
		DiscardPile: discardPile,
		Eliminated:  eliminated,
	}
}
