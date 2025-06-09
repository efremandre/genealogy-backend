package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/efremandre/genealogy-backend/internal/models"
	"github.com/efremandre/genealogy-backend/internal/storage"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	for _, u := range storage.Users {
		if u.Email == email {
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		}
	}

	user := models.User{
		ID:       int64(len(storage.Users) + 1),
		Email:    email,
		Password: string(hash),
	}

	storage.Users = append(storage.Users, user)
	err = storage.SaveUsers("users.json", storage.Users)
	if err != nil {
		log.Printf("Ошибка при сохранении пользователей: %v", err)
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowes", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" || password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	for _, u := range storage.Users {
		if u.Email == email {
			err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

			if err == nil {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("Login saccessful"))
				return
			}
		}
	}

	http.Error(w, "Invalid Email or Password", http.StatusUnauthorized)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowes", http.StatusBadRequest)
		return
	}

	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "Email required", http.StatusBadRequest)
		return
	}

	for _, u := range storage.Users {
		if u.Email == email {
			resp := struct {
				ID    int64  `json:"id"`
				Email string `json:"email"`
			}{
				ID:    u.ID,
				Email: u.Email,
			}

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(resp)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowes", http.StatusBadRequest)
		return
	}

	var resp []struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	}

	for _, u := range storage.Users {
		resp = append(resp, struct {
			ID    int64  `json:"id"`
			Email string `json:"email"`
		}{
			ID:    u.ID,
			Email: u.Email,
		})
	}

	if len(resp) == 0 {
		http.Error(w, "Users not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
