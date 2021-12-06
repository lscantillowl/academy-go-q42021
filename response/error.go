package response

import (
	"net/http"

	"github.com/lscantillowl/academy-go-q42021/domain/model"
)

func HandleError(code int, msg string) model.Response {
	return model.Response{
		Message: "Bad type parameter",
		Code:    http.StatusBadRequest,
	}
}
