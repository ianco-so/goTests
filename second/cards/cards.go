package cards

//var deckSize int = 52

var Suits = []string{"Diamonds", "Hearts", "Clubs", "Spades"}
var Ranks = []string{"Ace", "Two", "Three", "Four",
	"Five", "Six", "Seven", "Eight",
	"Nine", "Ten", "Jack", "Queen",
	"King"}

func Card(rank int, suit int) string {
	return Ranks[rank] + " of " + Suits[suit]
}
