package hand

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
  "github.com/jtschelling/deck/students/jtschelling/pkg/hand"
)

func TestDraw(t *testing.T) {
  deck := deck.New([]string{}, 0)
  hand := hand.Draw(deck, hand, 1)

  if len(hand.Cards) != 1 {
    t.Errorf("Did not draw 1 card.")
  }
}
