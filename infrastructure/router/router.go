package router

import (
	"log"
	"net/http"

	"github.com/lscantillowl/academy-go-q42021/interface/controller"

	"github.com/gorilla/mux"
)

// App struct to hold the router
type App struct {
	Router *mux.Router
}

// Initialize the app function to initialize the router
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

// Run the app function to start the server
func (a *App) Run(addr string) {
	log.Println("Server running...")
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		return
	}
}
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", controller.HomeHandler).Methods("GET")
	a.Router.HandleFunc("/characters", controller.GetCharacters).Methods("GET")
	a.Router.HandleFunc("/save_characters", controller.SaveCharacters).Methods("GET")
	a.Router.HandleFunc("/read_characters/{type}/{items}/{itemsPerWorker}", controller.ReadCharacters).Methods("GET")
}
