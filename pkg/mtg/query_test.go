package mtg

import (
	"fmt"
	"testing"
)

// Test CardReqeust
func TestCardReqeust(t *testing.T) {
	// Single Request
	t.Run("Test CardReqeust", func(t *testing.T) {
		cardResponseList, err := CardRequest("Sol Ring", 1)
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 1
		if got != want {
			t.Errorf("got %d cards want %d", got, want)
		}
	})

	// Multi Request
	t.Run("Test Bad CardReqeust", func(t *testing.T) {
		cardResponseList, err := CardRequest("Omnath", 3)
		if err != nil {
			t.Error(err)
		}
		got := len(*cardResponseList)
		want := 3
		if got != want {
			t.Errorf("got %d cards want %d", got, want)
		}
		for _, card := range *cardResponseList {
			fmt.Println(card)
			fmt.Println()
		}
	})
}
