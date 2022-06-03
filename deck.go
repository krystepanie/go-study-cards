package main

import (
	"fmt"
	ioutil "io/ioutil"
	"strings"
)

type Deck []string

const Seperator = ", "

func NewDeck() Deck {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := Deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d Deck) Print() {
	fmt.Print("Deck: ")
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d Deck) ToString() string {
	return strings.Join(d, Seperator)
}

func FromString(s string) Deck {
	return strings.Split(s, Seperator)
}

func (d Deck) Deal(handSize int) (hand Deck, remainingCards Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) SaveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.ToString()), 0666)
}

func ReadFromFile(filename string) (Deck, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read file %w", err)
	}
	s := string(file)
	return FromString(s), nil
}
