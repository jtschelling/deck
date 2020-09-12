package deck

import (
  "math/rand"
  "time"
)

type Card struct {
  Value int
  Suit string
}

var cardTypes = [13]string{
  "Ace",
  "2",
  "3",
  "4",
  "5",
  "6",
  "7",
  "8",
  "9",
  "10",
  "Jack",
  "Queen",
  "King",
}

var StandardDeckSize = 52

var Suits = [4]string{
  "club",
  "diamond",
  "heart",
  "spade",
}

func New(numJokers int) []Card {
  deck := make([]Card, (StandardDeckSize + numJokers))

  cardPosition := 0

  for _, suit := range Suits {
    for i := 0; i <= 12; i++ {
      deck[cardPosition] = Card {
        Value: i,
        Suit: suit,
      }
      cardPosition++
    }
  }

  for i := 0; i < numJokers; i++ {
    deck[StandardDeckSize + i] = Card {
      Value: -1,
      Suit: "",
    }
  }

  return deck
}

func Shuffle(deck []Card, style string) []Card {
  switch style {
  case "random":
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(deck), func(i, j int) {
      deck[i], deck[j] = deck[j], deck[i]
    })
  }

  return deck
}
