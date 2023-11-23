package Cards

type Suit int
type Rank int

const (
	Spades Suit = iota
	Hearts
	Clubs
	Diamonds
)

const (
	Ace Rank = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit       Suit
	Rank       Rank
	IsRevealed bool
}

// change the default string representation of a suit
func (suit Suit) String() string {
	switch suit {
	case Spades:
		return "Spades"
	case Hearts:
		return "Hearts"
	case Clubs:
		return "Clubs"
	case Diamonds:
		return "Diamonds"
	default:
		return "Unknown"
	}
}

// change the default string representation of a rank
func (rank Rank) String() string {
	switch rank {
	case Ace:
		return "Ace"
	case Two:
		return "Two"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Six:
		return "Six"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return "Unknown"
	}
}

// change the default string representation of a card
func (card Card) String() string {
	return card.Rank.String() + " of " + card.Suit.String()
}
