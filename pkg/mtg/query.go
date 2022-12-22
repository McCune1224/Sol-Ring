package mtg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BASE_URL = "https://api.scryfall.com"

	SEARCH_URL = BASE_URL + "/cards/search"
)

func CardSearch(name string) (*[]Card, error) {
	// Create URL query params
	v := url.Values{}
	// Card to search for
	v.Set("q", name)

	// Append query params to URL with '?' prefix
	queryURL := SEARCH_URL + "?" + v.Encode()

	// Create HTTP GET request to queryURL
	resp, err := http.Get(queryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read JSON response body and get the bytes from the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cardResponse ScryFallCardQueryResponse
	err = json.Unmarshal(body, &cardResponse)
	if err != nil {
		return nil, err
	}

	return cardResponse.Cards, nil
}
