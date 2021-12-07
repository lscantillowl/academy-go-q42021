package model

// Structo to handle vars in request
type Vars struct {
	Items          int    `json:"items"`
	ItemsPerWorker int    `json:"itemsPerWorker"`
	Type           string `json:"type"`
}
