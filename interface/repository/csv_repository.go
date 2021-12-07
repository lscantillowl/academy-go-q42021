package repository

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"

	logger "github.com/sirupsen/logrus"
)

// Function that create a CSV file and save data
func CreateCSV(data model.ResponseApi, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}
	SaveCSVData(data, f)
}

// Function that save data in a CSV file
func SaveCSVData(data model.ResponseApi, f *os.File) {
	writeCSV := csv.NewWriter(f)
	defer writeCSV.Flush()
	for _, pokemon := range data.Results {
		s := strings.Split(pokemon.Url, "/")
		id := s[len(s)-2]
		record := []string{
			id,
			pokemon.Name,
			pokemon.Url,
		}
		if err := writeCSV.Write(record); err != nil {
			log.Fatalln("Error writing record to file", err)
		}
	}
}

// function that convert a csv file to a character pokemon list
func ReadCSV(fileName string) []model.Character {
	csvLines := OpenCSV(fileName)

	pokemonsList := make([]model.Character, 0)

	for _, line := range csvLines {
		idCharacter, _ := strconv.Atoi(line[0])
		character := model.Character{
			Id:   idCharacter,
			Name: line[1],
		}
		pokemonsList = append(pokemonsList, character)

	}
	return pokemonsList
}

// Function that open a CSV file and return a [][]string
func OpenCSV(file string) [][]string {
	csvFile, err := os.Open(config.CSVFileOpen)
	if err != nil {
		logger.Error("Error opening CSV: ", err)
	}
	logger.Debug("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		logger.Error("Error reading CSV: ", err)
	}
	return csvLines
}

// Function that retrieve data from Pokemons API
func CallPokemonsAPI(url string) []byte {
	result, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()

	bodyBytes, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	return bodyBytes
}
