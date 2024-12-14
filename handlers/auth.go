package handlers

import (
	"encoding/json"
	"fmt"
	"icomers/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = []models.User{
	{ID: 1, Username: "dummy", Password: "$2a$10$amS7Gwe2c7dTpzcO5Rf7PuuNpRM03eyhou8Eyt3HbhdLUVANQVKTy", Email: "dummy@dummy.com", CreatedAt: time.Now()}, // Password: "password123"
}

// Secret key used for signing JWTs (In production, use an environment variable)
var jwtSecretKey = []byte("secret_key")

func createToken(user models.User) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    fmt.Sprintf("%d", user.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	// hash the password first
	err := newUser.HashPassword()
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	newUser.ID = len(users) + 1
	newUser.CreatedAt = time.Now()
	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	_ = json.NewDecoder(r.Body).Decode(&loginUser)

	var foundUser models.User

	for _, user := range users {
		if user.Username == loginUser.Username {
			foundUser = user
			break
		}
	}

	if foundUser.Username == "" || !foundUser.CheckPassword(loginUser.Password) {
		http.Error(w, "Username or password is invalid", http.StatusUnauthorized)
		return
	}

	token, err := createToken(foundUser)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "token invalid or expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
