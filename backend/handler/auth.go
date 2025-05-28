package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Deepankar20/shinra/backend/models"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserCred struct {
	Username string
	Email    string
	Password string
}

func AuthHandler(r chi.Router, db *gorm.DB) {
	h := Handler{DB: db}
	r.Post("/signin", h.signIn)
	r.Post("/signup", h.signUp)
	r.Post("/clerk-webhook", ClerkWebhook)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var creds UserCred

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user models.User

	token, err := generateJWT(user.ID, user.Email)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := h.DB.Table("users").Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
	w.Write([]byte("SignIn Success"))

}

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var creds UserCred

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Error Hashing Password", http.StatusInternalServerError)
	}

	user := models.User{
		Username: creds.Username,
		Email:    creds.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := h.DB.Table("users").Create(&user).Error; err != nil {
		http.Error(w, "Error Creating User", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))

}

func ClerkWebhook(w http.ResponseWriter, r *http.Request) {

}

var jwtKey = []byte("one-piece-is-real")

func generateJWT(userID uint, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
