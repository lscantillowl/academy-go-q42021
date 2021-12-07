package model

// Structure to parse the character data
type ResponseCharacter struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

// Structure to parse the response data
type ResponseApi struct {
	Count    int64               `json:"count"`
	Next     string              `json:"next"`
	Previous string              `json:"previous"`
	Results  []ResponseCharacter `json:"results"`
}
