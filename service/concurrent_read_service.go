package service

import (
	"strconv"

	"github.com/lscantillowl/academy-go-q42021/config"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
	"github.com/lscantillowl/academy-go-q42021/interface/repository"

	log "github.com/sirupsen/logrus"
)

// Function to read the csv file and return a map of strings
func ConcurrentPool(vars model.Vars) map[int][]string {

	items := vars.Items
	itemsPerWorker := vars.ItemsPerWorker

	jobs := make(chan []string, items)
	results := make(chan []string, items)

	for i := 0; i <= (items / itemsPerWorker); i++ {
		go Worker(i, jobs, results, vars.Type)
	}

	csvLines := repository.OpenCSV(config.CSVFileOpen)

	for index, line := range csvLines {

		jobs <- line

		if index >= items {
			break
		}
	}

	close(jobs)

	var result = make(map[int][]string)

	for a := 1; a <= items; a++ {

		t := <-results

		if t != nil {
			result[a] = t
		}
		log.Debug("Results: ", t)
	}

	log.Debug(result)

	return result
}

func Worker(id int, jobs <-chan []string, results chan<- []string, oddOrEven string) {
	for j := range jobs {
		log.Debug("Worker: ", id, " started  job", j)
		results <- TypeOfItem(j, oddOrEven)
		log.Debug("Worker: ", id, " finished job", j)
	}
}

func TypeOfItem(item []string, divisibility string) []string {

	switch divisibility {
	case "odd":
		id, err := strconv.Atoi(item[0])
		if err != nil {
			log.Error("There was a problem converting string to int: ", err)
		}

		if id%2 == 1 {
			return item
		} else {
			return nil
		}

	case "even":
		id, err := strconv.Atoi(item[0])
		if err != nil {
			log.Error("There was a problem converting string to int: ", err)
		}

		if id%2 == 0 {
			return item
		} else {
			return nil
		}

	default:
		log.Error("Bad type parameter")
		return nil
	}
}
