package main

import "fmt"

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Printf("something went wrong %s", err)
		}
	}()

	cards := NewDeck()
	cards.Print()
	fmt.Println("saving deck to a file")
	err = cards.SaveToFile("my_cards")

	fmt.Println("read deck from file")
	cardsFromFile, err := ReadFromFile("my_cards")
	cardsFromFile.Print()
	//hand, remainingCards := cards.Deal(5)
	//fmt.Printf("remainingCards: %s", remainingCards.ToString())
	//fmt.Printf("hand: %s", hand.ToString())

}
