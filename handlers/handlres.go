package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/cbuelvasc/twittor/middlewares"
	"github.com/cbuelvasc/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", middlewares.CheckDataBase(routers.Register)).Methods("POST")
	router.HandleFunc("/api/login", middlewares.CheckDataBase(routers.Login)).Methods("POST")
	router.HandleFunc("/api/profile", middlewares.CheckDataBase(middlewares.ValidateJWT(routers.Profile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
