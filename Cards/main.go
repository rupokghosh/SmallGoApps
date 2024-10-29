package main

import "fmt"

func main() {
	fmt.Println("Cards!")
	
	cards:= newDeck()
	cards.print()
}
