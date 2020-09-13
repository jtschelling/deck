package hand

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

// Tests deck.Draw() on a standard deck. Card should be added to hand and removed from deck.
func TestDraw(t *testing.T) {
  d := deck.New([]string{}, 0)
  h := New()
  d, h = Draw(d, h, 1)

  if len(h.Cards) != 1 {
    t.Errorf("Did not add drawn card to hand.")
  }

  if len(d.Cards) != 51 {
    t.Errorf("Did not remove drawn card from deck.")
  }
}

// Tests deck.Show() on a standard deck. Stdout should receive contents of hand.
func TestShow(t *testing.T) {
  d := deck.New([]string{}, 0)
  h := New()
  d, h = Draw(d, h, 5)
}

// Tests deck.Discard() on a standard deck. Correct card should be discarded and
// correct card should be added to the deck's discard pile.
func TestDiscard(t *testing.T) {
  d := deck.New([]string{}, 0)
  h := New()
  d, h = Draw(d, h, 2)
  expectedCardInHand := h.Cards[1]
  d, h = Discard(d, h, h.Cards[0])

  if len(h.Cards) > 1 {
    t.Errorf("Too few cards were removed from hand.")
  } else if len(h.Cards) < 1 {
    t.Errorf("Too many cards were removed from hand.")
  }

  if (h.Cards[0].Value != expectedCardInHand.Value || h.Cards[0].Suit != expectedCardInHand.Suit) {
    t.Errorf("Discarded incorrect card.")
  }

  if len(d.Discard) != 1 {
    t.Errorf("Card was not added to discard pile.")
  }
}
