package deck

import (
	"fmt"
	"testing"
)

//Sorting

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Spade})
	fmt.Println(Card{Rank: Nine, Suit: Diamond})
	fmt.Println(Card{Rank: Seven, Suit: Clove})
	fmt.Println(Card{Suit: Joker})

	//Output:
	//Ace of Hearts
	//Two of Spades
	//Nine of Diamonds
	//Seven of Cloves
	//Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	def := Card{Rank: Ace, Suit: Spade}
	if cards[0] != def {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

// Customizing sorting of the cards
func TestSort(t *testing.T) {
	cards := New(Sort(LessCards))
	def := Card{Rank: Ace, Suit: Spade}
	if cards[0] != def {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}
