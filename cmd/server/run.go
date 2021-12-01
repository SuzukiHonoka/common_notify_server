package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	collectAPI "nfly/api/collect"
	feedAPI "nfly/api/feed"
	pushAPI "nfly/api/push"
	userAPI "nfly/api/user"
	confServer "nfly/config/server"
	initialize "nfly/init/database"
)

func Run() {
	// opt: read extra file as config
	// todo-1: initialize, connect database and read data (user, config)
	initialize.DataBase()
	// todo-2: load plugin and workers
	// todo-3: start scheduler and timer or cronjob
	// todo-4: start the safeguard
	// todo-5: start the API server
	router := mux.NewRouter()
	router.HandleFunc("/login", userAPI.UserLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", userAPI.UserRegister).Methods(http.MethodPost)
	router.HandleFunc("/logout", userAPI.UserLogout).Methods(http.MethodGet)
	router.HandleFunc("/delete/{user}", userAPI.UserDelete).Methods(http.MethodDelete)
	router.HandleFunc("/push", pushAPI.Push).Methods(http.MethodPost)
	router.HandleFunc("/feeds", feedAPI.GetFeeds).Methods(http.MethodGet)
	router.HandleFunc("/collect/{uuid}", collectAPI.Collect).Methods(http.MethodPut)
	log.Println("NFLY API will listen tcp incomes at", confServer.Binds)
	log.Fatal(http.ListenAndServe(confServer.Binds, router))
}
