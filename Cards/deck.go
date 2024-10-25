package main

import "fmt"

type deck []string

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, c := range cardSuits {

		for _, v := range cardValues {
			cards = append(cards, c + " of " + v)
		}
	}
	return cards
}
