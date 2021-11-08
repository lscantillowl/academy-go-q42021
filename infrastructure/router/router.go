package router

import (
	"encoding/csv"
	"encoding/json"
	"fmt"

	"os"
	"strconv"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lscantillowl/academy-go-q42021/domain/model"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Println("Server running...")
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		return
	}
}
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.HomeHandler).Methods("GET")
	a.Router.HandleFunc("/characters", a.GetCharacters).Methods("GET")
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Welcome to the Q4 GO Bootcamp API"
	respondWithJSON(w, http.StatusOK, resp)
}

func (a *App) GetCharacters(w http.ResponseWriter, r *http.Request) {

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
