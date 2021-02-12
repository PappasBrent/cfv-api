package models

type Set struct {
	ID   uint64 `json:"-"`
	Name string `json:"Name"`
}

func (Set) TableName() string {
	return "set"
}
