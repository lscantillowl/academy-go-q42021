package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lscantillowl/academy-go-q42021/response"
	"github.com/lscantillowl/academy-go-q42021/service"
	"github.com/lscantillowl/academy-go-q42021/utils"
	logger "github.com/sirupsen/logrus"
)

func ReadCharacters(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	logger.Debug("Received as params Type: ", vars["type"], ", Items: ", vars["items"], ", itemsPerWorker: ", vars["itemsPerWorker"])

	if vars["type"] == "odd" || vars["type"] == "even" {
		result := service.ConcurrentPool(vars)
		response := response.HandleResponse(http.StatusOK, "Success", result)
		utils.RespondWithJSON(w, http.StatusOK, response)
	} else {
		response := response.HandleError(http.StatusBadRequest, "Bad type parameter")
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}
