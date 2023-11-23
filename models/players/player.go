package player

import (
	Card "github.com/ArragonElessar/PokerMania/models/cards"
)

type Player struct {
	FirstName     string
	LastName      string
	Money         int
	TablePosition int
	HoleCards     []Card.Card
}
