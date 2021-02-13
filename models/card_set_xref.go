package models

type CardSetXref struct {
	ID     uint64 `json:"-"`
	CardID uint64 `json:"-" gorm:"column:card_id"`
	SetID  uint64 `json:"-" gorm:"column:set_id"`
}

func (CardSetXref) TableName() string {
	return "card_set_xref"
}