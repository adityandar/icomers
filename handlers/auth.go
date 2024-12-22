package handlers

import (
	"encoding/json"
	"fmt"
	"icomers/database"
	"icomers/dto"
	"icomers/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
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
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	// hash the password first
	if err := newUser.HashPassword(); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		http.Error(w, "Error creating new user", http.StatusInternalServerError)
		return
	}

	response := dto.ConvertToUserResponse(newUser)

	w.Header().Set("Content-Type", "application/json")
	// TODO(adityandar): when returning the data, it still send the `password` field
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User
	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
	}

	var foundUser models.User

	if err := database.DB.Where("username = ? ", loginUser.Username).First(&foundUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Username or password is invalid", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		}
		return
	}

	if !foundUser.CheckPassword(loginUser.Password) {
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
