package models

type TournamentStatus struct {
	ID     uint64 `json:"-"`
	CardId uint64 `json:"-"`
	En     string `json:"En"`
	Jp     string `json:"Jp"`
	Kr     string `json:"Kr"`
	Th     string `json:"Th"`
	It     string `json:"It"`
}

func (TournamentStatus) TableName() string {
	return "tournament_status"
}
