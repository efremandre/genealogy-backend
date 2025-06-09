package main

import (
	"log"
	"net/http"

	"github.com/efremandre/genealogy-backend/internal/handlers"
	"github.com/efremandre/genealogy-backend/internal/storage"
)

func main() {
	var err error
	users, err := storage.LoadUsers("users.json")
	if err != nil {
		log.Fatalf("Не удалось загрузить пользователей: %v", err)
	}
	storage.Users = users
	log.Printf("Загружено пользователей: %d", len(storage.Users))
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/users", handlers.GetAllUsersHandler)
	http.HandleFunc("/user", handlers.GetUserHandler)

	http.HandleFunc("/relation/create", handlers.CreateRelativeHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
