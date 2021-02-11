package models

type Set struct {
	ID   uint64
	Name string
}

func (Set) TableName() string {
	return "set"
}
