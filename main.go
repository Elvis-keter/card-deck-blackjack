//go:generate stringer -type=Suit,Rank

package deck

import (
	"fmt"
	"sort"
)

type Suit uint8

const (
	Spade Suit = iota
	Diamond
	Clove
	Heart
	Joker
)

var suits = [...]Suit{Spade, Diamond, Clove, Heart}

type Rank uint8

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

type Card struct {
	Suit
	Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New(options ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank++ {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, options := range options {
		cards = options(cards)
	}
	//every suit and rank
	//add card(suit,rank)
	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, LessCards(cards))
	return cards
}

func Sort(LessCards func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, LessCards(cards))
		return cards
	}
}

func LessCards(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return sortA(cards[i]) < sortA(cards[j])
	}
}

func sortA(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}
