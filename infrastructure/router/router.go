package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lscantillowl/academy-go-q42021/interface/controller"
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
	a.Router.HandleFunc("/", controller.HomeHandler).Methods("GET")
	a.Router.HandleFunc("/characters", controller.GetCharacters).Methods("GET")
}
