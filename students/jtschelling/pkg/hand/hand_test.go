package hand

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

func TestDraw(t *testing.T) {
  deck := deck.New([]string{}, 0)
  hand := New()
  deck, hand = Draw(deck, hand, 1)

  if len(hand.Cards) != 1 {
    t.Errorf("Did not add drawn card to hand.")
  }

  if len(deck.Cards) != 51 {
    t.Errorf("Did not remove drawn card from deck.")
  }
}
