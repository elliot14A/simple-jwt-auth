package users

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/elliot14A/practice/database"
	"github.com/elliot14A/practice/models"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Id    uint
	Email string
	jwt.StandardClaims
}

func login(w http.ResponseWriter, r *http.Request) {
	req := request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := models.User{}

	err = database.DB.Model(&user).Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})

	secret := []byte(viper.GetString("JWT_SECRET"))
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Println("error signing jwt")
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token": tokenString,
	}

	writeResponse(w, response, http.StatusOK)
	return
}
