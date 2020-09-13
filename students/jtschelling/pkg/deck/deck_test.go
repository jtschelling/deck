package deck

import (
  "testing"
)

func TestCardTypeExcluded(t *testing.T) {
  shouldBeExcluded := "Ace"
  if !CardTypeExcluded([]string{shouldBeExcluded}, "Ace") {
    t.Errorf("%s should have been excluded", shouldBeExcluded)
  }

  if CardTypeExcluded([]string{}, "King") {
    t.Errorf("%s should not have been excluded", "King")
  }
}

func TestCreateCardTypes(t *testing.T) {
  cardTypesRemoved := []string{
    "Ace",
  }

  cardTypes := CreateCardTypes(cardTypesRemoved, false)
  if len(cardTypes) != 12 {
    t.Errorf("card not removed")
  }
}

func testUnsortedDeck(testDeck Deck, deckSize int, numJokers int) bool {
  expectedSuit  := 0
  expectedValue := 0

  for _, card := range testDeck.Cards {
    if (card.Suit != Suits[expectedSuit] || card.Value != CardTypes[card.Value].Value) {
      return false
    }
    expectedValue = (expectedValue + 1) % ((deckSize - numJokers) / len(Suits))

    // move to next suit if expectedValue wrapped around to 0
    if expectedValue == 0 {
      expectedSuit++

      // if jokers are in the deck we need to break when all the standard suits
      // have been exhausted
      if expectedSuit == len(Suits) {
        break
      }
    }
  }

  if numJokers > 0 {
    for i := 0; i < numJokers; i++ {
      if (testDeck.Cards[(deckSize - numJokers) + i].Suit != "" || testDeck.Cards[(deckSize - numJokers) + i].Value != -1) {
        return false
      }
    }
  }

  return true
}

func TestNew(t *testing.T) {
  numJokers := 0
  newDeck := New([]string{}, numJokers)

  if len(newDeck.Cards) != StandardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", StandardDeckSize + numJokers)
  }

  if testUnsortedDeck(newDeck, len(newDeck.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestNewWithRemoved(t *testing.T) {
  numJokers := 0
  newDeck := New([]string{}, numJokers)

  if len(newDeck.Cards) != StandardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", StandardDeckSize + numJokers)
  }

  if testUnsortedDeck(newDeck, len(newDeck.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestCreateEuchreDeck(t *testing.T) {
  numJokers := 1
  newDeck := New([]string{"2", "3", "4", "5", "6", "7", "8"}, numJokers)

  if len(newDeck.Cards) != (25) {
    t.Errorf("Deck is not %d cards", 25)
  }

  if testUnsortedDeck(newDeck, len(newDeck.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestNewWithJokers(t *testing.T) {
  numJokers := 5
  newDeck := New([]string{}, numJokers)

  if len(newDeck.Cards) != (StandardDeckSize + numJokers) {
    t.Errorf("Deck is not %d cards", (StandardDeckSize + numJokers))
  }

  if testUnsortedDeck(newDeck, len(newDeck.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestRandomShuffle(t *testing.T) {
  newDeck      := New([]string{}, 0)

  shuffled     := false
  orderedSuit  := 0
  orderedValue := 0

  // there is an infinitesimally small chance that the deck randomly sorts
  // in the same order as created, but I'm not going to account for that.
  shuffledDeck := Shuffle(newDeck, "random")
  for _, card := range shuffledDeck.Cards {
    if (card.Suit != Suits[orderedSuit] || card.Value != orderedValue) {
      shuffled = true
      break
    } else {
      shuffled = false
    }


    orderedValue = (orderedValue + 1) % 13

    // move to next suit if expectedValue wrapped around to 0
    if orderedValue == 0 {
      orderedSuit++
    }
  }

  if shuffled != true {
    t.Errorf("Deck was not shuffled randomly.")
  }
}
