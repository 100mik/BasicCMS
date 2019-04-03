package app

import "github.com/gorilla/mux"

func (mycms *App) LoadRoutes() {
	mycms.router = mux.NewRouter()

	mycms.router.HandleFunc("/hello", mycms.hello).Methods("GET")
}