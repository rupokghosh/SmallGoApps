package main

import "fmt"

func main() {
	fmt.Println("Cards!")
	
	cards:= newDeck()
	cards.shuffle()
	cards.print()
}
