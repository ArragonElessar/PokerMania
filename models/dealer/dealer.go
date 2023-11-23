package Dealer

import (
	Cards "github.com/ArragonElessar/PokerMania/models/cards"
	Deck "github.com/ArragonElessar/PokerMania/models/deck"
	Player "github.com/ArragonElessar/PokerMania/models/player"
)

type Dealer struct {
	Deck           *Deck.Deck
	PotMoney       int
	Name           string
	CommunityCards []Cards.Card
	Players        []Player.Player
}
