package deck

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

func TestCardTypeExcluded(t *testing.T) {
  shouldBeExcluded := "Ace"
  if !deck.CardTypeExcluded([]string{shouldBeExcluded}, "Ace") {
    t.Errorf("%s should have been excluded", shouldBeExcluded)
  }

  if deck.CardTypeExcluded([]string{}, "King") {
    t.Errorf("%s should not have been excluded", "King")
  }
}

func TestCreateCardTypes(t *testing.T) {
  cardTypesRemoved := []string{
    "Ace",
  }

  cardTypes := deck.CreateCardTypes(cardTypesRemoved, false)
  if len(cardTypes) != 12 {
    t.Errorf("card not removed")
  }
}

func testSortedDeck(testDeck []deck.Card, deckSize int, numJokers int) bool {
  expectedSuit  := 0
  expectedValue := 0

  for _, card := range testDeck {
    if (card.Suit != deck.Suits[expectedSuit] || card.Value != deck.CardTypes[card.Value].Value) {
      return false
    }
    expectedValue = (expectedValue + 1) % ((deckSize - numJokers) / len(deck.Suits))

    // move to next suit if expectedValue wrapped around to 0
    if expectedValue == 0 {
      expectedSuit++

      // if jokers are in the deck we need to break when all the standard suits
      // have been exhausted
      if expectedSuit == len(deck.Suits) {
        break
      }
    }
  }

  if numJokers > 0 {
    for i := 0; i < numJokers; i++ {
      if (testDeck[(deckSize - numJokers) + i].Suit != "" || testDeck[(deckSize - numJokers) + i].Value != -1) {
        return false
      }
    }
  }

  return true
}

func TestNew(t *testing.T) {
  numJokers := 0
  newDeck := deck.New([]string{}, numJokers)

  if len(newDeck) != deck.StandardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", deck.StandardDeckSize + numJokers)
  }

  if testSortedDeck(newDeck, len(newDeck), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestNewWithRemoved(t *testing.T) {
  numJokers := 0
  newDeck := deck.New([]string{}, numJokers)

  if len(newDeck) != deck.StandardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", deck.StandardDeckSize + numJokers)
  }

  if testSortedDeck(newDeck, len(newDeck), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestCreateEuchreDeck(t *testing.T) {
  numJokers := 1
  newDeck := deck.New([]string{"2", "3", "4", "5", "6", "7", "8"}, numJokers)

  if len(newDeck) != (25) {
    t.Errorf("Deck is not %d cards", 25)
  }

  if testSortedDeck(newDeck, len(newDeck), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestNewWithJokers(t *testing.T) {
  numJokers := 5
  newDeck := deck.New([]string{}, numJokers)

  if len(newDeck) != (deck.StandardDeckSize + numJokers) {
    t.Errorf("Deck is not %d cards", (deck.StandardDeckSize + numJokers))
  }

  if testSortedDeck(newDeck, len(newDeck), numJokers) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestRandomShuffle(t *testing.T) {
  newDeck      := deck.New([]string{}, 0)

  shuffled     := false
  orderedSuit  := 0
  orderedValue := 0

  // there is an infinitesimally small chance that the deck randomly sorts
  // in the same order as created, but I'm not going to account for that.
  shuffledDeck := deck.Shuffle(newDeck, "random")
  for _, card := range shuffledDeck {
    if (card.Suit != deck.Suits[orderedSuit] || card.Value != orderedValue) {
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
