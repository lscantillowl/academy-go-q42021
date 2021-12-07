package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
	"github.com/lscantillowl/academy-go-q42021/response"
	"github.com/lscantillowl/academy-go-q42021/service"
	"github.com/lscantillowl/academy-go-q42021/utils"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

// Home handler to test the api
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	response := response.HandleResponse(http.StatusOK, "Welcome to the Q4 GO Bootcamp API", nil)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Function to get characters from csv file and return it as json
func GetCharacters(w http.ResponseWriter, r *http.Request) {
	pokemonsList := service.ReadCSV(config.CSVFileOpen)
	response := response.HandleResponse(http.StatusOK, "Success", pokemonsList)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Save characters to csv file from api call
func SaveCharacters(w http.ResponseWriter, r *http.Request) {
	bodyBytes := service.CallPokemonsAPI(config.ApiURL)
	var responseObject model.ResponseApi
	json.Unmarshal(bodyBytes, &responseObject)
	service.CreateCSV(responseObject, config.CSVFileSave)
	response := response.HandleResponse(http.StatusOK, "Successfully saved characters to csv file", nil)
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// function that reads all characters from the csv file using the service concurrent pool(Worker pool)
func ReadCharacters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	items, err := strconv.Atoi(vars["items"])
	if err != nil {
		logger.Error("There was an error converting items parameter to int: ", err)
	}
	itemsPerWorker, err := strconv.Atoi(vars["itemsPerWorker"])
	if err != nil {
		logger.Error("There was an error converting items parameter to int: ", err)
	}
	logger.Debug("Received as params Type: ", vars["type"], ", Items: ", vars["items"], ", itemsPerWorker: ", vars["itemsPerWorker"])
	params := model.Vars{
		Type:           vars["type"],
		Items:          items,
		ItemsPerWorker: itemsPerWorker,
	}
	if params.Type == "odd" || params.Type == "even" {
		result := service.ConcurrentPool(params)
		response := response.HandleResponse(http.StatusOK, "Success", result)
		utils.RespondWithJSON(w, http.StatusOK, response)
	} else {
		response := response.HandleError(http.StatusBadRequest, "Bad type parameter")
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
	}
}
