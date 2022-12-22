package mtg

import (
	"testing"
)

// Test CardReqeust
func TestCardReqeust(t *testing.T) {
	// Creature Request
	t.Run("Creature", func(t *testing.T) {
		cardResponseList, err := CardSearch("Omnath, Locus of Rage")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Artifact Request
	t.Run("Artifact", func(t *testing.T) {
		cardResponseList, err := CardSearch("Sol Ring")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Planeswalker Request
	t.Run("Planeswalker", func(t *testing.T) {
		cardResponseList, err := CardSearch("Narset, Parter of Veils")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Instant Request
	t.Run("Instant", func(t *testing.T) {
		cardResponseList, err := CardSearch("Counterspell")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Socery Request
	t.Run("Socery", func(t *testing.T) {
		cardResponseList, err := CardSearch("Brainstorm")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Vague Request
	t.Run("Ambiguous", func(t *testing.T) {
		cardResponseList, err := CardSearch("Teysa")
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got < want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})
}
