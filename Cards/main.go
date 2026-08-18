package main

import "fmt"

func main() {

	// var card string = "Ace of Spades"
	// card = "new value for card"
	cards := newDeck()
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println([]byte("\nRemaining deck\n"))
	remainingDeck.print()

	// hand.saveToFile("myCards.txt")

	// deck := newDeckFromFile("myCards.txt")
	// deck.print()

	hand.print()
    hand.shuffle()
	hand.print()

}
