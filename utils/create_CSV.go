package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
)

func CreateCSV(data model.ResponseApi) {
	f, err := os.Create(config.CSVFileSave)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

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
