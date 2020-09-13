// Impements operation on cards in a deck
// Shuffle
// Draw
// Discard
package deck

import (
  "math/rand"
  "time"
)

/////////////
// STRUCTS //
/////////////

type Card struct {
  Value int
  Suit string
}

type CardType struct {
  name string
  Value int
}

type Deck struct {
  Cards []Card
  Discard []Card
}

/////////////
// GLOBALS //
/////////////

var cardTypes = []CardType{
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

var standardDeckSize = 52

var suits = [4]string{
  "club",
  "diamond",
  "heart",
  "spade",
}

//////////////////////////
// FUNCTIONS ON STRUCTS //
//////////////////////////

func (d Deck) Draw() (Deck, Card) {
  c := d.Cards[0]
  d = Deck {
    Cards: d.Cards[1:],
  }
  return d, c
}

////////////
// PUBLIC //
////////////

// Creates new deck.
// Arguments:
// removedCards - Array of card type names to remove from deck. Ex. {"Ace", "4", "King"}
// numJokers - Integer number of joker cards to add to the deck.
func New(removedCards []string, numJokers int) Deck {
  cardPosition := 0

  jokerIncluded := false
  if numJokers > 0 {
    jokerIncluded = true
  }
  includedCardTypes := createCardTypes(removedCards, jokerIncluded)

  numNonJockerCards := len(suits) * (len(includedCardTypes))
  if jokerIncluded {
    numNonJockerCards = len(suits) * (len(includedCardTypes) - 1)
  }

  d := make([]Card, (numNonJockerCards + numJokers))

  for _, suit := range suits {
    for _, cardType := range includedCardTypes {
      // add jokers in the next for loop because they have a special value and no suit
      if jokerIncluded && cardType.name == "Joker" {
        continue
      }

      d[cardPosition] = Card {
        Value: cardType.Value,
        Suit: suit,
      }
      cardPosition++
    }
  }

  for i := 0; i < numJokers; i++ {
    d[len(d) - numJokers + i] = Card {
      Value: -1,
      Suit: "",
    }
  }

  return Deck {
    Cards: d,
    Discard: make([]Card, 0),
  }
}

// Places provided card in deck's discard pile.
func AddToDiscard(d Deck, c Card) Deck {
  return Deck {
    Cards: d.Cards,
    Discard: append(d.Discard, c),
  }
}

// Shuffle the deck according to the specified style
// Argument s accepts the following:
// - "random"
func Shuffle(d Deck, s string) Deck {
  switch s {
  case "random":
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(d.Cards), func(i, j int) {
      d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    })
  }

  return d
}

/////////////
// PRIVATE //
/////////////

// Creates list of card types to include in a new deck. Iterates over a standard
// list and removes card types specified by the removedCards argument.
func createCardTypes(removedCards []string, jokerIncluded bool) []CardType {
  if (removedCards == nil && jokerIncluded == false) {
    return cardTypes[0:12]
  } else if removedCards == nil {
    return cardTypes
  }

  var includedCardTypes []CardType
  for i := 0; i < len(cardTypes); i++ {
    if !cardTypeExcluded(removedCards, cardTypes[i].name) {
      if cardTypes[i].name == "Joker" && !jokerIncluded {
        continue
      }

      includedCardTypes = append(includedCardTypes, cardTypes[i])
    }
  }

  return includedCardTypes
}

// Checks if the cardType argument matches any values in the removedCards list.
// Return true if the card type should be excluded from the final deck.
func cardTypeExcluded(removedCards []string, cardType string) bool {
  for _, cardName := range removedCards {
    if cardName == cardType {
      return true
    }
  }

  return false
}
