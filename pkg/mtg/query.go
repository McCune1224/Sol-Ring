package mtg

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	BASE_URL  = "https://api.magicthegathering.io/v1/"
	CARDS_URL = BASE_URL + "cards"
)

func CardRequest(name string, pageSize ...int) (*[]Card, error) {
	// Default to 1 if no pageSize is provided or if pageSize is over 100
	if len(pageSize) == 0 || pageSize[0] < 100 {
		pageSize = append(pageSize, 1)
	}

	// Create URL query params
	v := url.Values{}
	v.Set("name", name)
	strPageSize := strconv.Itoa(pageSize[0])
	v.Set("pageSize", strPageSize)

	// Append query params to URL with '?' prefix
	queryURL := CARDS_URL + "?" + v.Encode()

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

	var cardResponse CardResponse
	err = json.Unmarshal(body, &cardResponse)
	if err != nil {
		return nil, err
	}

	return cardResponse.Cards, nil
}
