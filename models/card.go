package models

type Card struct {
	ID                 uint64           `json:"-" gorm:"primaryKey"`
	CardType           string           `json:"CardType"`
	Clan               string           `json:"Clan"`
	Critical           int              `json:"Critical"`
	DesignIllus        string           `json:"DesignIllus"`
	Effect             string           `json:"Effect"`
	Flavor             string           `json:"Flavor"`
	Format             string           `json:"Format"`
	Grade              string           `json:"Grade"`
	Illust             string           `json:"Illust"`
	IllustColor        string           `json:"IllustColor"`
	Illust2            string           `json:"Illust2"`
	Illust3            string           `json:"Illust3"`
	Illust4            string           `json:"Illust4"`
	Illust5            string           `json:"Illust5"`
	ImageUrlEn         string           `json:"ImageUrlEn"`
	ImageUrlJp         string           `json:"ImageUrlJp"`
	ImaginaryGift      string           `json:"ImaginaryGift"`
	Italian            string           `json:"Italian"`
	Kana               string           `json:"Kana"`
	Kanji              string           `json:"Kanji"`
	Korean             string           `json:"Korean"`
	LimitationText     string           `json:"LimitationText"`
	MangaIllust        string           `json:"MangaIllust"`
	Name               string           `json:"Name" gorm:"not null"`
	Nation             string           `json:"Nation"`
	Note               string           `json:"Note"`
	OtherNames         string           `json:"OtherNames"`
	Phonetic           string           `json:"Phonetic"`
	Power              int              `json:"Power"`
	Race               string           `json:"Race"`
	RideSkill          string           `json:"RideSkill"`
	Sets               []Set            `json:"Sets" gorm:"many2many:card_set_xref"`
	TournamentStatuses TournamentStatus `json:"TournamentStatuses"`
	Shield             int              `json:"Shield"`
	Skill              string           `json:"Skill"`
	Thai               string           `json:"Thai"`
	Translation        string           `json:"Translation"`
	TriggerEffect      string           `json:"TriggerEffect"`
}

func (Card) TableName() string {
	return "card"
}
