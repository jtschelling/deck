// Implements operations on player's hands
// Play
// Draw
// Discard
package hand

import (
  "fmt"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

type Hand struct {
  Cards []deck.Card
}

func New() Hand {
  cards := []deck.Card {
    deck.Card {
      Value: 1,
      Suit: "heart",
    },
  }

  return Hand {
    Cards: cards,
  }
}

func Draw(d deck.Deck, h Hand, numToDraw int) (deck.Deck, Hand) {
  var drawnCard deck.Card
  for i := 0; i < numToDraw; i++ {
    d, drawnCard = d.Draw()
    h = addToHand(h, drawnCard)
  }

  return d, h
}

func Show(h Hand) {
  for _, card := range h.Cards {
    fmt.Println(card)
  }
}

func Discard(d deck.Deck, c deck.Card) Hand {
  return New()
}

func addToHand(h Hand, c deck.Card) Hand {
  return New()
}
