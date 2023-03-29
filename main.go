package main

type Suit struct{ Name, Symbol string }

// Facevalue will be 2 to 10, Jack Queen King ro Ace
type Facevalue struct {
	Name  string
	Value int
}

// Card is a Suit and Facevalue
type Card struct {
	Suit  Suit
	Value Facevalue
}

// GreaterThan tells if one card's Value is GreaterThan the other
func (card *Card) GreaterThan(b *Card) bool {
	return card.Value.Value > b.Value.Value
}

// LessThan tells if one card's Value is LessThan the other
func (card *Card) LessThan(b *Card) bool {
	return card.Value.Value < b.Value.Value
}

// Equal returns true or false based on the Facevalue
func (card *Card) Equal(b *Card) bool {
	return card.Value.Value == b.Value.Value
}

// Facecard returns true for J,Q,K,A, false for all others
func (card *Card) Facecard() (ans bool) {
	n := card.Value.Name
	return n == "Jack" || n == "Queen" || n == "King" || n == "Ace"
}
