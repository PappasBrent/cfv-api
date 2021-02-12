package models

type CardSetXref struct {
	ID     uint64 `json:"-"`
	CardId uint64 `json:"-"`
	SetId  uint64 `json:"-"`
}

func (CardSetXref) TableName() string {
	return "card_set_xref"
}
