package models

import (
	"cfv-api/constants"
	"encoding/json"
)

type Set struct {
	ID    uint64 `json:"-"`
	Name  string `json:"name"`
	Cards []Card `json:"cards" gorm:"many2many:card_set_xrefs"`
}

// TableName sets table name for gorm
func (Set) TableName() string {
	return constants.SetsTableName
}

// MarshalJSON marshals only the set name
func (s *Set) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Name)
}
