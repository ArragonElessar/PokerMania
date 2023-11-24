package Dealer

import (
	"fmt"

	Cards "github.com/ArragonElessar/PokerMania/models/cards"
	Deck "github.com/ArragonElessar/PokerMania/models/deck"
	Player "github.com/ArragonElessar/PokerMania/models/player"
)

const MaxPlayers = 8
const InitialBuyIn = 100
const BigBlindAmount = 2 // Dollars

type Dealer struct {
	Deck           *Deck.Deck
	PotMoney       int // TODO Case when there is an all in
	Name           string
	CommunityCards []Cards.Card
	ActivePlayers  []*Player.Player
	SmallBlind     int
	BigBlind       int
}

/*
1. Wait for 2 players
2. Create and shuffle a deck
3. Distribute cards in sequential order to each player
4. Take top 3 cards from deck and put into community cards
5. Burn a card
6. Take top card from deck and put into community cards
7. Burn a card
8. Take top card from deck and put into community cards
9. Game begins
	9.1 Game Rules and Gameplay
10. Game ends -> Return the pot money back to the winner
11. Check if any player has exited and return back to step 1
*/

// function to create a new dealer
func CreateNewDealer(name string, smallBlind int) *Dealer {

	// check that name is non-null
	if name == "" {
		fmt.Println("Cannot create a dealer with no name")
		return nil
	}

	deck := Deck.InitializeDeck()
	deck.ShuffleDeck()

	return &Dealer{Deck: deck,
		PotMoney:       0,
		Name:           name,
		CommunityCards: []Cards.Card{},
		ActivePlayers:  []*Player.Player{},
		SmallBlind:     smallBlind,
		BigBlind:       2 * smallBlind}
}

// helper function
func (dealer *Dealer) isPlayerActive(player *Player.Player) bool {

	for _, p := range dealer.ActivePlayers {
		if p.FirstName == player.FirstName && p.LastName == player.LastName {
			return true
		}
	}
	return false
}

// functions to add or remove the players
func (dealer *Dealer) AddPlayer(player *Player.Player) int {

	// check if there is any space for the player
	if len(dealer.ActivePlayers) == MaxPlayers {
		fmt.Println("Cannot add player, max players reached")
		return -1
	}
	// check if this player already exists in the game
	if dealer.isPlayerActive(player) {
		fmt.Println("Player is already playing")
		return -2
	}
	// add the player to the game
	dealer.ActivePlayers = append(dealer.ActivePlayers, player)
	return 1

}

func (dealer *Dealer) RemovePlayer(player *Player.Player) {

	// remove if the player exists in the game
	for i, _ := range dealer.ActivePlayers {
		if dealer.ActivePlayers[i] == player {
			dealer.ActivePlayers = append(dealer.ActivePlayers[:i], dealer.ActivePlayers[i+1:]...)
			return
		}
	}
}

// function to check if game can be started
func (dealer *Dealer) CanStartGame() bool {

	// check if there are atleast 2 players
	if len(dealer.ActivePlayers) < 2 {
		return false
	}
	// check if all the players have enough money for initial buy in, else give warning to remove the player
	var allPlayersReady bool = true
	for _, player := range dealer.ActivePlayers {
		if player.Money < BigBlindAmount {
			fmt.Println("Player", player.FirstName, player.LastName, "does not have enough money to play.")
			fmt.Println("Either get a buy in or please exit")
			allPlayersReady = false
		}
	}

	if !allPlayersReady {
		return false
	}

	// everything okay, start game.
	return true
}

// function to award a buy in to a player
func (dealer *Dealer) AwardBuyIn(player *Player.Player, amount int) bool {
	// check if the player exists in the game
	if !dealer.isPlayerActive(player) {
		fmt.Println("Player is not playing")
		return false
	}

	player.Money += amount
	fmt.Println("Player", player.FirstName, player.LastName, "has been awarded a buy in of", amount)
	fmt.Println("Player new balance is", player.Money)

	return true
}

// function to distribute cards to a players / dealer
func (dealer *Dealer) DistributeCards() bool {

	// check if game can be started
	if !dealer.CanStartGame() {
		return false
	}

	// if yes, create a new deck and shuffle it
	dealer.Deck.ShuffleDeck()

	// deal two cards to each player in sequential order
	for i := 0; i < 2; i++ {
		for _, player := range dealer.ActivePlayers {
			player.HoleCards = append(player.HoleCards, dealer.Deck.DealTopCard())
		}
	}
	// burn a card
	dealer.Deck.BurnCard()
	// add 3 cards to community cards
	for i := 0; i < 3; i++ {
		dealer.CommunityCards = append(dealer.CommunityCards, dealer.Deck.DealTopCard())
	}
	// burn a card
	dealer.Deck.BurnCard()

	// add 1 card to community cards
	dealer.CommunityCards = append(dealer.CommunityCards, dealer.Deck.DealTopCard())

	return true
}

// function to see revealed community cards
func (dealer *Dealer) PrintCommunityCards() {
	fmt.Println("Community Cards are:")
	for _, card := range dealer.CommunityCards {
		if card.IsRevealed {
			fmt.Println(card.Rank, "of", card.Suit)
		} else {
			fmt.Println("Hidden")
		}
	}
}
