package mtg

type Rulings []struct {
	Date string `json:"date"`
	Text string `json:"text"`
}

type ForeignNames []struct {
	Name         string      `json:"name"`
	Text         string      `json:"text"`
	Type         string      `json:"type"`
	Flavor       interface{} `json:"flavor"`
	ImageURL     string      `json:"imageUrl"`
	Language     string      `json:"language"`
	Multiverseid int         `json:"multiverseid"`
}

type Legalities []struct {
	Format   string `json:"format"`
	Legality string `json:"legality"`
}

// Card Represents the content structure of a 200 response MTG API '/card' endpoint
type Card struct {
	Name          string       `json:"name"`
	ManaCost      string       `json:"manaCost,omitempty"`
	Cmc           float64      `json:"cmc"`
	Colors        []string     `json:"colors,omitempty"`
	ColorIdentity []string     `json:"colorIdentity,omitempty"`
	Type          string       `json:"type"`
	Supertypes    []string     `json:"supertypes,omitempty"`
	Types         []string     `json:"types"`
	Subtypes      []string     `json:"subtypes,omitempty"`
	Rarity        string       `json:"rarity"`
	Set           string       `json:"set"`
	SetName       string       `json:"setName"`
	Text          string       `json:"text"`
	Artist        string       `json:"artist"`
	Number        string       `json:"number"`
	Power         string       `json:"power,omitempty"`
	Toughness     string       `json:"toughness,omitempty"`
	Layout        string       `json:"layout"`
	Multiverseid  string       `json:"multiverseid,omitempty"`
	ImageURL      string       `json:"imageUrl,omitempty"`
	Rulings       Rulings      `json:"rulings,omitempty"`
	ForeignNames  ForeignNames `json:"foreignNames,omitempty"`
	Printings     []string     `json:"printings"`
	OriginalText  string       `json:"originalText,omitempty"`
	OriginalType  string       `json:"originalType,omitempty"`
	Legalities    Legalities   `json:"legalities,omitempty"`
	ID            string       `json:"id"`
	Watermark     string       `json:"watermark,omitempty"`
	Flavor        string       `json:"flavor,omitempty"`
	Hand          string       `json:"hand,omitempty"`
	Life          string       `json:"life,omitempty"`
	Variations    []string     `json:"variations,omitempty"`
}

type CardResponse struct {
	Cards *[]Card `json:"cards"`
}

// Response from MTG API is always a list of card object, even if only 1 card is queried
