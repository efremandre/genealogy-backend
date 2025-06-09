package models

type Relative struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Relation  string `json:"relation"`
}
