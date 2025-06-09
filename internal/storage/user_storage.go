package storage

import (
	"encoding/json"
	"os"

	"github.com/efremandre/genealogy-backend/internal/models"
)

var Users []models.User

func LoadUsers(filename string) ([]models.User, error) {
	file, err := os.Open(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []models.User{}, nil
		}
		return nil, err
	}
	defer file.Close()
	var users []models.User
	err = json.NewDecoder(file).Decode(&users)
	return users, err
}

func SaveUsers(filename string, users []models.User) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(users)
}
