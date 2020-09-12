package deck

import (
  "testing"
  "github.com/jtschelling/deck/students/jtschelling/pkg/deck"
)

func testSortedDeck(testDeck []deck.Card) bool {
  expectedSuit  := 0
  expectedValue := 0

  for _, card := range testDeck {
    if (card.Suit != deck.Suits[expectedSuit] || card.Value != expectedValue) {
      return false
    }
    expectedValue = (expectedValue + 1) % 13

    // move to next suit if expectedValue wrapped around to 0
    if expectedValue == 0 {
      expectedSuit++

      // if jokers are in the deck we need to break when all the standard suits
      // have been exhausted
      if expectedSuit == 4 {
        break
      }
    }
  }

  return true
}

func TestNew(t *testing.T) {
  numJokers := 0
  newDeck := deck.New(numJokers)

  if len(newDeck) != deck.StandardDeckSize + numJokers {
    t.Errorf("Deck is not %d cards", deck.StandardDeckSize + numJokers)
  }

  if testSortedDeck(newDeck) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }
}

func TestNewWithJokers(t *testing.T) {
  numJokers := 5
  newDeck := deck.New(numJokers)

  if len(newDeck) != (deck.StandardDeckSize + numJokers) {
    t.Errorf("Deck is not %d cards", (deck.StandardDeckSize + numJokers))
  }

  if testSortedDeck(newDeck) == false {
    t.Errorf("New unsorted deck is not in the correct order.")
  }

  for i := 0; i < numJokers; i++ {
    if (newDeck[deck.StandardDeckSize + i].Suit != "" || newDeck[deck.StandardDeckSize + i].Value != -1) {
      t.Errorf("New unsorted deck does not contain jokers or are in the incorrect position.")
      t.FailNow()
    }
  }
}

func TestRandomShuffle(t *testing.T) {
  newDeck      := deck.New(0)

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
