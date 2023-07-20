package users

import (
	"net/http"
)

func protected(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value("claims").(*Claims)
	writeResponse(w, claims, http.StatusOK)
}
