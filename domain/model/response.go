package model

type Response struct {
	Message  string      `json:"message"`
	Code     int         `json:"code"`
	Response interface{} `json:"response"`
}
