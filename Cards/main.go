package main

import "fmt"

func main() {
	fmt.Println("Cards!")
	d1 := newDeck()
	d2, d3 := deal(d1, 5)
	fmt.Println(d2)
	fmt.Println(d3)
}
