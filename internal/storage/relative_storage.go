package storage

import (
	"encoding/json"
	"os"

	"github.com/efremandre/genealogy-backend/internal/models"
)

var Relatives []models.Relative

func LoadRelatives(filename string) ([]models.Relative, error) {
	file, err := os.Open(filename)

	if err != nil {
		if os.IsNotExist(err) {
			return []models.Relative{}, nil
		}
		return nil, err
	}
	defer file.Close()
	var relative []models.Relative
	err = json.NewDecoder(file).Decode(&relative)
	return relative, err
}

func SaveRelatives(filename string, relative []models.Relative) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(relative)
}
