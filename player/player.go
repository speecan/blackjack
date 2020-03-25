package player

import "github.com/speecan/blackjack/game"

type (
	Player interface {
		GetName() string
		CallHit(hand []*game.Card) bool
	}
)
