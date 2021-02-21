package models

import "cfv-api/constants"

type TournamentStatus struct {
	ID     uint64 `json:"-"`
	CardID uint64 `json:"-"`
	En     string `json:"en"`
	Jp     string `json:"jp"`
	Kr     string `json:"kr"`
	Th     string `json:"th"`
	It     string `json:"it"`
}

// TableName sets table name for gorm
func (TournamentStatus) TableName() string {
	return constants.TournamentStatusesTableName
}
