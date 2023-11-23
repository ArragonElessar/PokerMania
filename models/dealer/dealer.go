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
	ActivePlayers  []Player.Player
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

// helper function
func (dealer *Dealer) isPlayerActive(player *Player.Player) bool {

	for _, p := range dealer.ActivePlayers {
		if p == player {
			return true
		}
	}
	return false
}

// functions to add or remove the players
func (dealer *Dealer) AddPlayer(player *Player.Player) bool {

	// check if there is any space for the player
	if len(dealer.ActivePlayers) == MaxPlayers {
		fmt.Println("Cannot add player, max players reached")
		return false
	}
	// check if this player already exists in the game
	if dealer.isPlayerActive(player) {
		fmt.Println("Player is already playing")
		return false
	}
	// add the player to the game
	dealer.ActivePlayers = append(dealer.ActivePlayers, player)
	return true

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
