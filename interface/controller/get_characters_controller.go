package controller

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
	"github.com/lscantillowl/academy-go-q42021/response"
	"github.com/lscantillowl/academy-go-q42021/utils"
)

// Function to get characters from csv file and return it as json
func GetCharacters(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open(config.CSVFileOpen)
	if err != nil {
		log.Fatalf("can't open characters file - %s", err)
	}
	defer file.Close()

	csvLines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	pokemonsList := make([]model.Character, 0)

	for _, line := range csvLines {
		idCharacter, _ := strconv.Atoi(line[0])
		character := model.Character{
			Id:   idCharacter,
			Name: line[1],
		}
		pokemonsList = append(pokemonsList, character)

	}
	response := response.HandleResponse(http.StatusOK, "Success", pokemonsList)
	utils.RespondWithJSON(w, http.StatusOK, response)
}
