package handler

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func UserRoutes(r chi.Router, db *gorm.DB) {
	h := Handler{DB: db}
	r.Get("/", h.GetAllUsers)
	r.Get("/{id}", h.GetUserById)
	r.Post("/", h.CreateUser)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get all users"))
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(fmt.Sprintf("User ID : %s\n", id)))

}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}
