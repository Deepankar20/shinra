package router

import (
	"github.com/Deepankar20/shinra/backend/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *chi.Mux {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/user", func(r chi.Router) {
		handler.UserRoutes(r, db)
	})

	r.Route("/auth", func(r chi.Router) {
		handler.AuthHandler(r, db)
	})

	return r
}
