package service

import (
	"github.com/lscantillowl/academy-go-q42021/domain/model"
	"github.com/lscantillowl/academy-go-q42021/interface/repository"
)

// Service to read the csv file with a filename
func ReadCSV(fileName string) []model.Character {
	return repository.ReadCSV(fileName)
}

// Service to create a csv file sending data and filename
func CreateCSV(data model.ResponseApi, fileName string) {
	repository.CreateCSV(data, fileName)
}

// Service to get data from pokemon api
func CallPokemonsAPI(url string) []byte {
	return repository.CallPokemonsAPI(url)
}
