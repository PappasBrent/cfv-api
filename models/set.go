package models

import (
	"cfv-api/constants"
	"encoding/json"
)

type Set struct {
	ID    uint64 `json:"-"`
	Name  string `json:"Name"`
	Cards []Card `json:"Cards" gorm:"many2many:card_set_xrefs"`
}

func (Set) TableName() string {
	return constants.SetsTableName
}

func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name)
}
