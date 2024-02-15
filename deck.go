package main

import "fmt"

//create a new type of 'deck'

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSutis := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jacks", "Queen", "King"}

	for i := 0; i < len(cardSutis); i++ {
		for j := 0; j < len(cardValue); j++ {
			cards = append(cards, cardSutis[i]+" of "+cardValue[j])
		}
	}
	return cards
}
