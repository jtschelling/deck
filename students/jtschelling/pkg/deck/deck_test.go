package deck

import (
  "testing"
)

// Tests cardTypeExcluded() in various combinations of card types and removal lists.
func TestCardTypeExcluded(t *testing.T) {
  shouldBeExcluded := "Ace"
  if !cardTypeExcluded([]string{shouldBeExcluded}, "Ace") {
    t.Errorf("%s should have been excluded", shouldBeExcluded)
  }

  shouldBeExcluded = "2"
  if !cardTypeExcluded([]string{shouldBeExcluded}, "2") {
    t.Errorf("%s should have been excluded", shouldBeExcluded)
  }

  if cardTypeExcluded([]string{}, "King") {
    t.Errorf("%s should not have been excluded", "King")
  }
}

// Tests createCardTypes() with different removal lists and expects a list of
// card types excluding types from the removal list
func TestCreateCardTypes(t *testing.T) {
  cardTypes := createCardTypes([]string{}, false)
  if len(cardTypes) != 13 {
    t.Errorf("Card type removed when remove array argument is empty.")
  }

  cardTypesRemoved := []string{
    "Ace",
  }

  cardTypes = createCardTypes(cardTypesRemoved, false)
  if len(cardTypes) != 12 {
    t.Errorf("Ace not removed.")
  }
}

// Tests deck properties of a standard euchre decks when a removal list and 1 joker
// is provided to the New() function.
func TestCreateEuchreDeck(t *testing.T) {
  numJokers := 1
  d := New([]string{"2", "3", "4", "5", "6", "7", "8"}, numJokers)

  if len(d.Cards) != (25) {
    t.Errorf("Deck is not %d cards", 25)
  }

  if testUnsortedDeck(d, len(d.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

// Tests New() with standard settings. Should provide an unsorted 52 card deck
func TestNew(t *testing.T) {
  numJokers := 0
  d := New([]string{}, numJokers)

  if len(d.Cards) != standardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", standardDeckSize + numJokers)
  }

  if testUnsortedDeck(d, len(d.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

// Tests New() with jokers added. Should return an unsorted deck that contains
// 52 + (number of jokers specified) cards
func TestNewWithJokers(t *testing.T) {
  numJokers := 5
  d := New([]string{}, numJokers)

  if len(d.Cards) != (standardDeckSize + numJokers) {
    t.Errorf("Deck is not %d cards", (standardDeckSize + numJokers))
  }

  if testUnsortedDeck(d, len(d.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

// Tests New() with a list of card types to remove from the deck and expects
// a deck without cards in the removal list
func TestNewWithRemoved(t *testing.T) {
  numJokers := 0
  d := New([]string{}, numJokers)

  if len(d.Cards) != standardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", standardDeckSize + numJokers)
  }

  if testUnsortedDeck(d, len(d.Cards), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

// Tests Shuffle() with "random" style specified. Expects a deck in a random order.
func TestRandomShuffle(t *testing.T) {
  d := New([]string{}, 0)

  shuffled     := false
  orderedSuit  := 0
  orderedValue := 0

  // there is an infinitesimally small chance that the deck randomly sorts
  // in the same order as created, but I'm not going to account for that.
  shuffledDeck := Shuffle(d, "random")
  for _, c := range shuffledDeck.Cards {
    if (c.Suit != suits[orderedSuit] || c.Value != orderedValue) {
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

/////////////
// PRIVATE //
/////////////

// Private function to test for an unsorted deck, a few tests want to check this easily.
func testUnsortedDeck(d Deck, deckSize int, numJokers int) bool {
  expectedSuit  := 0
  expectedValue := 0

  for _, c := range d.Cards {
    if (c.Suit != suits[expectedSuit] || c.Value != cardTypes[c.Value].Value) {
      return false
    }
    expectedValue = (expectedValue + 1) % ((deckSize - numJokers) / len(suits))

    // move to next suit if expectedValue wrapped around to 0
    if expectedValue == 0 {
      expectedSuit++

      // if jokers are in the deck we need to break when all the standard suits
      // have been exhausted
      if expectedSuit == len(suits) {
        break
      }
    }
  }

  if numJokers > 0 {
    for i := 0; i < numJokers; i++ {
      if (d.Cards[(deckSize - numJokers) + i].Suit != "" || d.Cards[(deckSize - numJokers) + i].Value != -1) {
        return false
      }
    }
  }

  return true
}
