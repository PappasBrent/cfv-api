package models

type CardSetXref struct {
	ID     uint64
	CardId uint64
	En     string
	Jp     string
	Kr     string
	Th     string
	It     string
}

func (CardSetXref) TableName() string {
	return "card_set_xref"
}
