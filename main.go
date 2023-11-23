package main

import (
	Deck "github.com/ArragonElessar/PokerMania/models/deck"
)

func main() {

	deck := Deck.InitializeDeck()
	deck.ShuffleDeck()
	deck.PrintDeck()
}
