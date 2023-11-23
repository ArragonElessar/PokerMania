package Deck

import (
	"fmt"
	"math/rand"

	Cards "github.com/ArragonElessar/PokerMania/models/cards"
)

type Deck struct {
	cards []Cards.Card
}

// function to create a new deck
func InitializeDeck() *Deck {

	// create a new deck
	var deck *Deck
	var cards []Cards.Card
	deck = &Deck{cards}

	// for all 4 suits, for 13 ranks, create a card and append it to the deck
	for i := 0; i < 4; i++ { // TODO 4
		for j := 1; j <= 13; j++ {
			var card Cards.Card
			card.Suit = Cards.Suit(i)
			card.Rank = Cards.Rank(j)
			deck.cards = append(deck.cards, card)
		}
	}
	fmt.Println("Successfully created a new deck")

	return deck
}

// function to print the deck
func (deck *Deck) PrintDeck() {
	for i := 0; i < len(deck.cards); i++ {
		fmt.Println(deck.cards[i].Rank, "of", deck.cards[i].Suit)
	}
	fmt.Println("")
}

// function to shuffle the deck
func (deck *Deck) ShuffleDeck() {

	// check if the deck has 52 cards
	if len(deck.cards) != 52 {
		deck = InitializeDeck()
	}

	// shuffle the deck
	rand.Shuffle(len(deck.cards), func(i, j int) { deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i] })

}

// function to remove the top card of the deck
func (deck *Deck) DealTopCard() Cards.Card {
	var dealtCard Cards.Card
	dealtCard, deck.cards = deck.cards[0], deck.cards[1:]
	return dealtCard
}

// function to burn a card
func (deck *Deck) BurnCard() {

	topCard := deck.DealTopCard()
	deck.cards = append(deck.cards, topCard)

}
