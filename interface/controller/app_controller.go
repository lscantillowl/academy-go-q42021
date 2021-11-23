package controller

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lscantillowl/academy-go-q42021/domain/model"
)

// respondWithJSON write json response format to the response writer
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Can not convert payload to JSON - %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Home handler to test the api
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Welcome to the Q4 GO Bootcamp API"
	respondWithJSON(w, http.StatusOK, resp)
}

// Function to get characters from csv file and return it as json
func GetCharacters(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("characters.csv")
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
	respondWithJSON(w, http.StatusOK, pokemonsList)
}

// Save characters to csv file from api call
func SaveCharacters(w http.ResponseWriter, r *http.Request) {
	const myUrl = "https://pokeapi.co/api/v2/pokemon/"
	response, err := http.Get(myUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.ResponseApi
	json.Unmarshal(bodyBytes, &responseObject)

	f, err := os.Create("pokemons.csv")
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	writeCSV := csv.NewWriter(f)
	defer writeCSV.Flush()

	for _, pokemon := range responseObject.Results {
		s := strings.Split(pokemon.Url, "/")
		id := s[len(s)-2]
		record := []string{
			id,
			pokemon.Name,
			pokemon.Url,
		}
		if err := writeCSV.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	respondWithJSON(w, http.StatusOK, "Success")
}
