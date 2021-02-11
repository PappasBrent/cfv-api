package models

type TournamentStatus struct {
	ID     uint64
	CardId uint64
	SetId  uint64
}

func (TournamentStatus) TableName() string {
	return "tournament_status"
}
