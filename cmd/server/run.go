package server

import (
	collectAPI "common_notify_server/api/collect"
	feedAPI "common_notify_server/api/feed"
	pushAPI "common_notify_server/api/push"
	userAPI "common_notify_server/api/user"
	confServer "common_notify_server/config/server"
	initialize "common_notify_server/init/database"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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
	router.HandleFunc("/push", pushAPI.Push).Methods(http.MethodPost)
	router.HandleFunc("/feeds", feedAPI.GetFeeds).Methods(http.MethodGet)
	router.HandleFunc("/collect/{uuid}", collectAPI.Collect).Methods(http.MethodPut)
	log.Println("NFLY API will listen tcp incomes at", confServer.Binds)
	log.Fatal(http.ListenAndServe(confServer.Binds, router))
}
