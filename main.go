package main

import (
	"log"
	"net/http"

	"github.com/elliot14A/practice/database"
	"github.com/elliot14A/practice/handlers/users"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	err := database.Init()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	middleware.Logger(router)
	middleware.Recoverer(router)
	router.Mount("/users", users.InitRoutes())
	log.Println("server started at port 8080")
	http.ListenAndServe(":8080", router)
}
