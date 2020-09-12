package deck

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

func New(t *testing.T) {
  newDeck := deck.New()

  if len(newDeck) != 52 {
    t.Errorf("Deck is not 52 cards")
  }
}
