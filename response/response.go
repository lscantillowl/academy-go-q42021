package response

import "github.com/lscantillowl/academy-go-q42021/domain/model"

func HandleResponse(code int, msg string, data interface{}) model.Response {
	return model.Response{
		Code:     code,
		Message:  msg,
		Response: data,
	}
}
