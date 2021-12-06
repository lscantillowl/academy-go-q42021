package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
	"github.com/lscantillowl/academy-go-q42021/response"
	"github.com/lscantillowl/academy-go-q42021/utils"
)

// Save characters to csv file from api call
func SaveCharacters(w http.ResponseWriter, r *http.Request) {
	const myUrl = config.ApiURL
	result, err := http.Get(myUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()

	bodyBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.ResponseApi
	json.Unmarshal(bodyBytes, &responseObject)

	utils.CreateCSV(responseObject)
	response := response.HandleResponse(http.StatusOK, "Successfully saved characters to csv file", nil)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
