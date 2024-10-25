package main

import "fmt"

func main() {
	fmt.Println("Cards!")
	d1 := newDeck()
	d2 := d1.toString()
	fmt.Println(d2)
	d1.saveToFile("mycards")
}
