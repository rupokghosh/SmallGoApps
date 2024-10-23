package main

/*
1. Read CSV File line by line
2. Creating a map that maps questions as keys to answers as values
3. Iterate through the map to display the keys to users
4. Lookup the solution using the map
5. If user solution matches answer, increment score variable
6. Display total problems ( size of hashmap ) and value of score
7.
*/
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Quiz Game")

	file, err := os.Open("records.csv")
	if err != nil {
		log.Fatal("error reading file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error", err)
	}
	
	recordMap = make(map[string] int)
	for _, record:= range records{
		fmt.Println(record)
	}
	
}
