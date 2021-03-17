package models

import (
	"cfv-api/constants"
	"encoding/json"
)

type Card struct {
	ID                 uint64           `json:"id" gorm:"primaryKey"`
	CardType           string           `json:"cardtype"`
	Clan               string           `json:"clan"`
	Critical           int              `json:"critical"`
	DesignIllus        string           `json:"designillus"`
	Effect             string           `json:"effect"`
	Flavor             string           `json:"flavor"`
	Format             string           `json:"format"`
	Grade              string           `json:"grade"`
	Illust             string           `json:"illust"`
	IllustColor        string           `json:"illustcolor"`
	Illust2            string           `json:"illust2"`
	Illust3            string           `json:"illust3"`
	Illust4            string           `json:"illust4"`
	Illust5            string           `json:"illust5"`
	ImageURLEn         string           `json:"imageurlen"`
	ImageURLJp         string           `json:"imageurljp"`
	ImaginaryGift      string           `json:"imaginarygift"`
	Italian            string           `json:"italian"`
	Kana               string           `json:"kana"`
	Kanji              string           `json:"kanji"`
	Korean             string           `json:"korean"`
	LimitationText     string           `json:"limitationtext"`
	MangaIllust        string           `json:"mangaillust"`
	Name               string           `json:"name" gorm:"not null"`
	Nation             string           `json:"nation"`
	Note               string           `json:"note"`
	OtherNames         string           `json:"othernames"`
	Phonetic           string           `json:"phonetic"`
	Power              int              `json:"power"`
	Race               string           `json:"race"`
	RideSkill          string           `json:"rideskill"`
	Sets               []Set            `json:"sets" gorm:"many2many:card_set_xrefs"`
	TournamentStatuses TournamentStatus `json:"tournamentstatuses"`
	Shield             int              `json:"shield"`
	Skill              string           `json:"skill"`
	Thai               string           `json:"thai"`
	Translation        string           `json:"translation"`
	TriggerEffect      string           `json:"triggereffect"`
}

// TableName sets table name for gorm
func (Card) TableName() string {
	return constants.CardsTableName
}

// MarshalJSON comples the card image urls
func (c *Card) MarshalJSON() ([]byte, error) {
	// See http://choly.ca/post/go-json-marshalling/
	type Alias Card
	result := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	return json.Marshal(result)
}
