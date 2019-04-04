package app

import "github.com/gorilla/mux"

func (mycms *App) LoadRoutes() {
	mycms.router = mux.NewRouter()

	mycms.router.HandleFunc("/hello", mycms.hello).Methods("GET")
	mycms.router.HandleFunc("/register", mycms.register).Methods("POST")
	mycms.router.HandleFunc("/login", mycms.login).Methods("POST")
	mycms.router.HandleFunc("/logout", mycms.logout).Methods("GET")
	mycms.router.HandleFunc("/getPosts", mycms.getPosts).Methods("GET")
	mycms.router.HandleFunc("/addPost", mycms.addPost).Methods("POST")
	}