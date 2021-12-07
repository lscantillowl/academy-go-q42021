package response

import "github.com/lscantillowl/academy-go-q42021/domain/model"

// Function that return a model.Response to parse data to json

func HandleResponse(code int, msg string, data interface{}) model.Response {
	return model.Response{
		Code:     code,
		Message:  msg,
		Response: data,
	}
}
