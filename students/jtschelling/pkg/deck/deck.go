package deck

// import (
//
// )

type Card struct {
  value Int
  suit String
}

var suits = []string{
  "club",
  "diamond",
  "heart",
  "spade",
}

func New() []Card (
  deck := make([]Card, 52)

  for _, suit := range suits {
    for i, i <= 13, i++ {
      card :=
      deck = append(deck, Card {
        value: i,
        suit: suit,
      })
    }
  }

  return deck
)
