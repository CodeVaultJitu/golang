package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Clubs", "hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four" }

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) joinString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.joinString()), 0666)
}

func newReadFile (filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPos := 	r.Intn(len(d)-1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}