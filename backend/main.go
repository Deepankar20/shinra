package main

import (
	"log"
	"net/http"

	"github.com/Deepankar20/shinra/backend/db"
	"github.com/Deepankar20/shinra/backend/router"
)

func main() {
	database, _ := db.InitDB()

	r := router.NewRouter(database)

	log.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
