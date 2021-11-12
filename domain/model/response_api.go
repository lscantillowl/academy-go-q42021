package model

type Response struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type ResponseApi struct {
	Count    int64      `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Response `json:"results"`
}
