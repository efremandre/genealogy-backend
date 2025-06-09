package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/efremandre/genealogy-backend/internal/models"
	"github.com/efremandre/genealogy-backend/internal/storage"
)

func CreateRelativeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowes", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	birthDate := r.FormValue("birth_date")
	relationType := r.FormValue("relation")

	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if len(name) > 100 {
		http.Error(w, "Name too long", http.StatusBadRequest)
	}

	if birthDate == "" {
		http.Error(w, "BirthDate  is required", http.StatusBadRequest)
		return
	}

	if relationType == "" {
		http.Error(w, "Relation is required", http.StatusBadRequest)
		return
	}

	relative := models.Relative{
		ID:        int64(len(storage.Relatives) + 1),
		Name:      name,
		BirthDate: birthDate,
		Relation:  relationType,
	}

	storage.Relatives = append(storage.Relatives, relative)
	err := storage.SaveRelatives("relatives.json", storage.Relatives)
	if err != nil {
		log.Printf("Ошибка при сохранении пользователей: %v", err)
		http.Error(w, "Failed to save user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(relative)

}
