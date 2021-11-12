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

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Welcome to the Q4 GO Bootcamp API"
	respondWithJSON(w, http.StatusOK, resp)
}

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

func SaveCharacters(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
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
