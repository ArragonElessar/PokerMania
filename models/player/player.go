package player

import (
	"fmt"

	Card "github.com/ArragonElessar/PokerMania/models/cards"
)

type Move int

const (
	Check Move = iota
	Bet
	Call
	Raise
	Fold
	AllIn
)

type Player struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Money         int    `json:"money"`
	ChipsValue    int    `json:"chipsvalue"`
	TablePosition int    `json:"tableposition"`
	HoleCards     []Card.Card
}

/*
1. Player can ask dealer to join a table
2. Player can ask for a buy in
3. Game Starts
	3.1 Cards are distributed
	3.2 Player can see their cards
	3.3 Player can make a move
	3.4 Player can see the community cards, when revealed
	3.5 Player will see what others have bet
	3.6 Similarly till the game ends
4. Game Ends, player can choose to continue / leave
*/

// function to create a new player
func CreateNewPlayer(firstName string, lastName string, money int) *Player {

	// check that at least one of firstname or lastname is non-null
	if firstName == "" && lastName == "" {
		fmt.Println("Cannot create a player with no name")
		return nil
	}

	return &Player{firstName, lastName, money, -1, -1, []Card.Card{}}
}
