package main

import (
	"fmt"
	"third/deck"
)

func main() {
	cards := deck.New()
	cards.Print()
	fmt.Println("There are", len(cards), "cards in the deck.")
}
