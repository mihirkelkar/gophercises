package deck

import (
	"math/rand"
	"time"
)

//Card : represents a complete card in a deck
type Card struct {
	value string
	suite string
}

//Deck : Returns a complete deck
func Deck() []Card {
	var deck []Card
	cardValues := []string{"ace", "two", "three", "four", "five", "six", "seven", "eight",
		"nine", "ten", "jack", "queen", "king"}

	cardSuites := []string{"jack", "hearts", "spades", "clubs"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			deck = append(deck, Card{value: value, suite: suite})
		}
	}
	return deck
}

//ShuffleDeck : Shuffles the deck and return a value
func ShuffleDeck(d []Card) []Card {
	for index := range d {
		rand.Seed(time.Now().UTC().UnixNano())
		newIndex := rand.Intn(52)
		d[index], d[newIndex] = d[newIndex], d[index]
	}
	return d
}

//AddJokers : Adds a specified number of jokers to the deck
func AddJokers(num int, d []Card) []Card {
	for num > 0 {
		d = append(d, Card{value: "Joker", suite: "Joker"})
	}
	return d
}

//Filter : Specify things to remove
func Filter(suite []string, values []string, d []Card) []Card {
	var newDeck []Card
	// filter out suites that we don't want
	for _, card := range d {
		for _, suite := range suite {
			if card.suite != suite {
				newDeck = append(newDeck, card)
			}
		}
	}

	//filter out numbers from the filtered suite deck
	j := 0
	for _, card := range newDeck {
		for _, value := range values {
			if card.value != value {
				newDeck[j] = card
				j++
			}
		}
	}
	return newDeck[:j]
}
