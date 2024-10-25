package main

import "fmt"

func main() {
	fmt.Println("Cards!")
	
	new := newDeckFromFile("mycards")
	new.print()
}
