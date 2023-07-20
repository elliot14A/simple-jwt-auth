package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func InitRoutes() chi.Router {
	router := chi.NewRouter()
	router.Post("/register", register)
	router.Post("/login", login)
	router.With(middleware).Get("/protected", protected)
	return router
}

func writeResponse(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonReponse, _ := json.Marshal(body)
	w.Write(jsonReponse)
}
