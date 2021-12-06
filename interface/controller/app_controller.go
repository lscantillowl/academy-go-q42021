package controller

import (
	"net/http"

	"github.com/lscantillowl/academy-go-q42021/utils"
)

// Home handler to test the api
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Welcome to the Q4 GO Bootcamp API"
	utils.RespondWithJSON(w, http.StatusOK, resp)
}
