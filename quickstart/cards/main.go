package main

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// // card = "Five of Diamonds"

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// // fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.print()
}
