package models

import (
	"encoding/json"
)

type Set struct {
	ID    uint64 `json:"-"`
	Name  string `json:"Name"`
	Cards []Card `json:"Cards" gorm:"many2many:card_set_xref"`
}

func (Set) TableName() string {
	return "set"
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name)
}
