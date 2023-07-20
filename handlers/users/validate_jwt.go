package users

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

func middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			log.Println("1")
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		claims, err := validateToken(token)
		if err != nil {
			log.Println("2")
			log.Println(err.Error())
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validateToken(jwtToken string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(jwtToken, claims, func(t *jwt.Token) (interface{}, error) {
		secret := viper.GetString("JWT_SECRET")
		return []byte(secret), nil
	})
	log.Println(token.Valid)
	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}
