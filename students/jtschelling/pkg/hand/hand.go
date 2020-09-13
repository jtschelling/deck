// Implements operations on player's hands
// Draw
// Discard
// Play
package hand

import (
  "fmt"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

/////////////
// STRUCTS //
/////////////

type Hand struct {
  Cards []deck.Card
}

////////////
// PUBLIC //
////////////

// Initialize
func New() Hand {
  return Hand {
    Cards: []deck.Card{},
  }
}

// Draws a card from the provided deck and adds it to a hand.
func Draw(d deck.Deck, h Hand, numToDraw int) (deck.Deck, Hand) {
  var drawnCard deck.Card
  for i := 0; i < numToDraw; i++ {
    d, drawnCard = d.Draw()
    h = addToHand(h, drawnCard)
  }

  return d, h
}

// Places a card from a hand into the deck's discard pile.
func Discard(d deck.Deck, h Hand, c deck.Card) (deck.Deck, Hand) {
  for ndx, card := range h.Cards {
    if card.Value == c.Value && card.Suit == c.Suit {
      h.Cards = append(h.Cards[:ndx], h.Cards[ndx+1:]...)
      break
    }
  }
  return deck.AddToDiscard(d, c), h
}

// Outputs cards in hand to stdout
func Show(h Hand) {
  for _, card := range h.Cards {
    fmt.Println(card)
  }
}

/////////////
// PRIVATE //
/////////////

// Adds card to hand struct Cards field
func addToHand(h Hand, c deck.Card) Hand {
  return Hand {
    Cards: append(h.Cards, c),
  }
}
