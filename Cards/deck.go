package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"math/rand"
)

type deck []string

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, c := range cardSuits {

		for _, v := range cardValues {
			cards = append(cards, c+" of "+v)
		}
	}
	return cards
}

func deal(d deck, dealSize int) (hand, other deck) {

	return d[:dealSize], d[dealSize:]

}

func (d deck) toString() string {
	res := []string(d)            // converting deck into an array of string
	return strings.Join(res, ",") // converting array of strings into individual strings separated by a comma

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {

	receivingDeck, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
	}
	s := strings.Split(string(receivingDeck), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // use this to generate a new int for every second
	r := rand.New(source)                           // use the source object to create the new random number
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
