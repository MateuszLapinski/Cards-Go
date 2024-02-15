package main

import "fmt"

func main() {
	//var card string = "Ace of Spades"
	//tworzenie zmiennej / nazwa zmiennej /typ zmiennej
	// card:= newCard()

	// fmt.Println(card)

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()

	for i := 1; i <= len(cards); i++ {
		fmt.Println(i, cards[i-1])
	}

	fmt.Println("------------------------------------------")

	cards = newDeck()

	cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
