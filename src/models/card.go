package models

type Card struct {
	ID             uint64
	CardType       string
	Clan           string
	Critical       int
	DesignIllus    string
	Effect         string
	Flavor         string
	Format         string
	Grade          string
	Illust         string
	IllustColor    string
	Illust2        string
	Illust3        string
	Illust4        string
	Illust5        string
	ImageUrlEn     string
	ImageUrlJp     string
	ImaginaryGift  string
	Italian        string
	Kana           string
	Kanji          string
	Korean         string
	LimitationText string
	MangaIllust    string
	Name           string // TODO: Add not null requirement
	Nation         string
	Note           string
	OtherNames     string
	Phonetic       string
	Power          int
	Race           string
	RideSkill      string
	Shield         int
	Skill          string
	Thai           string
	Translation    string
	TriggerEffect  string
}

func (Card) TableName() string {
	return "card"
}
