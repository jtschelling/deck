// Implements operations on player's hands
// Play
// Draw
// Discard
package hand

import (
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

type Hand struct {
  Cards []deck.Card
}

func Draw(deck deck.Deck, hand Hand, numToDraw int) Hand {
  for i := 0; i < numToDraw; i++ {
    drawnCard := deck.Draw(*deck)
    addToHand(*hand, drawnCard)
  }

  return hand
}

func Show(hand Hand) {
  for _, card := range hand.Cards {
    fmt.Print()
  }
}

func Discard(deck deck.Deck, card deck.Card) Hand {

}

func addToHand(hand *Hand, card deck.Card) {

}
