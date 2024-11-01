package game

type Team int

const (
	TeamWhite Team = iota
	TeamBlack
)

type Player struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Team        Team   `json:"team"`
	Supply      []Unit `json:"supply"`
	Hand        []Unit `json:"hand"`
	Bag         []Unit `json:"bag"`
	DiscardPile []Unit `json:"discardPile"`
}
