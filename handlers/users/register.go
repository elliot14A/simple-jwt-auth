package users

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"

	"github.com/elliot14A/practice/database"
	"github.com/elliot14A/practice/models"
	"golang.org/x/crypto/bcrypt"
)

func (r *request) Validate() error {
	email := r.Email
	password := r.Password

	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regEx := regexp.MustCompile(regexPattern)

	if !regEx.MatchString(email) {
		return errors.New("Invalid email")
	}

	if len(password) < 4 {
		return errors.New("password is too short")
	}
	return nil
}

func (r *request) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), err
}

func register(w http.ResponseWriter, r *http.Request) {
	req := &request{}
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := req.HashPassword()
	if err != nil {
		log.Println("error while hashing password")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
	}

	var count int64 = 0
	err = database.DB.Model(&models.User{}).Where("email = ?", user.Email).Count(&count).Error
	if err != nil {
		log.Println("error while creating the user")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if count > 0 {
		http.Error(w, "user with same exists", http.StatusConflict)
		return
	}

	err = database.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Println("error while creating the user")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"email": user.Email,
		"id":    user.Id,
	}

	json.NewEncoder(w).Encode(response)
	return
}
