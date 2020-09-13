// Impements operation on cards in a deck
// Shuffle
// Draw
// Discard
package deck

import (
  "math/rand"
  "time"
)

type Card struct {
  Value int
  Suit string
}

type CardType struct {
  name string
  Value int
}

var CardTypes = []CardType{
  CardType {
    name: "Ace",
    Value: 0,
  },
  CardType {
    name: "2",
    Value: 1,
  },
  CardType {
    name: "3",
    Value: 2,
  },
  CardType {
    name: "4",
    Value: 3,
  },
  CardType {
    name: "5",
    Value: 4,
  },
  CardType {
    name: "6",
    Value: 5,
  },
  CardType {
    name: "7",
    Value: 6,
  },
  CardType {
    name: "8",
    Value: 7,
  },
  CardType {
    name: "9",
    Value: 8,
  },
  CardType {
    name: "10",
    Value: 9,
  },
  CardType {
    name: "Jack",
    Value: 10,
  },
  CardType {
    name: "Queen",
    Value: 11,
  },
  CardType {
    name: "King",
    Value: 12,
  },
  CardType {
    name: "Joker",
    Value: -1,
  },
}

var StandardDeckSize = 52

var Suits = [4]string{
  "club",
  "diamond",
  "heart",
  "spade",
}

func CreateCardTypes(removedCards []string, jokerIncluded bool) []CardType {
  if (removedCards == nil && jokerIncluded == false) {
    return CardTypes[0:12]
  } else if removedCards == nil {
    return CardTypes
  }

  var includedCardTypes []CardType
  for i := 0; i < len(CardTypes); i++ {
    if !CardTypeExcluded(removedCards, CardTypes[i].name) {
      if CardTypes[i].name == "Joker" && !jokerIncluded {
        continue
      }

      includedCardTypes = append(includedCardTypes, CardTypes[i])
    }
  }

  return includedCardTypes
}

func CardTypeExcluded(removedCards []string, cardType string) bool {
  for _, cardName := range removedCards {
    if cardName == cardType {
      return true
    }
  }

  return false
}

func New(removedCards []string, numJokers int) []Card {
  cardPosition := 0

  jokerIncluded := false
  if numJokers > 0 {
    jokerIncluded = true
  }
  includedCardTypes := CreateCardTypes(removedCards, jokerIncluded)

  numNonJockerCards := len(Suits) * (len(includedCardTypes))
  if jokerIncluded {
    numNonJockerCards = len(Suits) * (len(includedCardTypes) - 1)
  }

  deck := make([]Card, (numNonJockerCards + numJokers))

  for _, suit := range Suits {
    for _, cardType := range includedCardTypes {
      // add jokers in the next for loop because they have a special value and no suit
      if jokerIncluded && cardType.name == "Joker" {
        continue
      }

      deck[cardPosition] = Card {
        Value: cardType.Value,
        Suit: suit,
      }
      cardPosition++
    }
  }

  for i := 0; i < numJokers; i++ {
    deck[len(deck) - numJokers + i] = Card {
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
