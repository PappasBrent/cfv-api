package v1

import (
	"cfv-api/constants"
	"cfv-api/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// There has to be a way to replace just the Sets member of the Card
// model to be an array string, but this works for now
type cardResponseBody struct {
	// The ID of the card to search for
	// example: 43
	ID uint64 `json:"id" gorm:"primaryKey"`

	// The type of the card
	// example: Trigger Unit
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	CardType string `json:"cardtype"`

	// The clan the card belongs to
	// example: Shadow Paladin
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Clan string `json:"clan"`

	// The card's critical value
	// example: 1
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Critical int `json:"critical"`

	// Card designer / illustrator
	// example: Azusa / 天城望
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	DesignIllus string `json:"designillus"`

	// The card's effect text
	// example: (You may only have up to four cards with "HEAL" in a deck.)
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Effect string `json:"effect"`

	// The card's flavor text
	// example: (V-TD04): Those with the will to fight will never give up!(V-BT04): I don't believe you wish to die like this!(V-BT06): Change the pain of your wounds into anger. And stand up once more!
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Flavor string `json:"flavor"`

	// The legal formats to play the card in
	// example: Standard / Premium Standard
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Format string `json:"format"`

	// The card's grade
	// example: 0
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Grade int `json:"grade"`

	// Additional card illustrator field #1
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Illust string `json:"illust"`

	// Additional card illustrator and colorer
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	IllustColor string `json:"illustcolor"`

	// Additional card illustrator field #2
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Illust2 string `json:"illust2"`

	// Additional card illustrator field #3
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Illust3 string `json:"illust3"`

	// Additional card illustrator field #4
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Illust4 string `json:"illust4"`

	// Additional card illustrator field #5
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Illust5 string `json:"illust5"`

	// A URL to a English scan of the card
	// example: http://cf-vanguard.cards/assets/card-images/43-en.png
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	ImageURLEn string `json:"imageurlen"`
	// TODO: Update this comment to match the live server

	// A URL to a Japanese scan of the card
	// example: http://cf-vanguard.cards/assets/card-images/43-jp.png
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	ImageURLJp string `json:"imageurljp"`

	// The card's imaginary gift
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	ImaginaryGift string `json:"imaginarygift"`

	// Italian translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Italian string `json:"italian"`

	// Kana translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Kana string `json:"kana"`

	// Kanji translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Kanji string `json:"kanji"`

	// Korean translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Korean string `json:"korean"`

	// The card's limitation text
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	LimitationText string `json:"limitationtext"`

	// The illustrator of the card in the manga
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	MangaIllust string `json:"mangaillust"`

	// The English name of the card
	// example: Abyss Healer
	Name string `json:"name" gorm:"not null"`

	// The nation the card belongs to
	// example: United Sanctuary
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Nation string `json:"nation"`

	// Any additional notes on the card form the wiki
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Note string `json:"note"`

	// Other for the card
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	OtherNames string `json:"othernames"`

	// Phonetic pronunciation of the card's name in Japanese
	// example: Abisu Hīrā
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Phonetic string `json:"phonetic"`

	// The card's power
	// example: 5000
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Power int `json:"power"`

	// The card's race
	// example: Angel
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Race string `json:"race"`

	// The card's ride skill
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	RideSkill string `json:"rideskill"`

	// The list of sets that the card belongs to
	// example: ["V Booster Set 04: Vilest! Deletor","V Trial Deck 04: Ren Suzugamori","V Extra Booster 12: Team Dragon's Vanity!","V Special Series 03: Start Deck Blaster Dark","V Booster Set 06: Phantasmal Steed Restoration"]
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Sets []string `json:"sets" gorm:"many2many:card_set_xrefs"`

	// The tournament statuses of the card
	// example: {"en":"Unrestricted","jp":"Unrestricted","kr":"Unrestricted","th":"Unrestricted","it":"Unrestricted"}
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	TournamentStatuses models.TournamentStatus `json:"tournamentstatuses"`

	// The card's shield value
	// example: 20000
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Shield int `json:"shield"`

	// The card's skill name
	// example: Boost
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Skill string `json:"skill"`

	// Thai translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Thai string `json:"thai"`

	// The literal translation of the card's name
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	Translation string `json:"translation"`

	// The card's trigger effect
	// example: Heal / +10000
	// Extensions:
	// ---
	// x-nullable: true
	// ---
	TriggerEffect string `json:"triggereffect"`
}

// swagger:route GET /card cards getCard
//
// Returns a single card with the specified ID
//
//	Responses:
// 		200: cardResponse

// A single card
// swagger:response cardResponse
type cardResponse struct {
	// A single card
	// in: Body
	Body cardResponseBody
}

// swagger:route GET /cards cards getCards
//
// Returns a list of cards matching the specified criteria
//
// 	Responses:
//		200: cardsResponse

// A list of cards
// swagger:response cardsResponse
type cardsResponse struct {
	// An array of cards
	// in: Body
	Body []cardResponseBody
}

// GetCard returns a single card as JSON
func GetCard(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	if id, err := strconv.Atoi(c.Query("id")); err != nil {
		c.JSON(http.StatusBadRequest, invalidCardIDError())
	} else {
		cardResult := models.Card{}

		if result := db.Preload("Sets").Preload("TournamentStatuses").
			Where("id = ?", id).
			Find(&cardResult); result.RowsAffected == 1 {
			c.JSON(200, cardResult)
		} else {
			c.JSON(http.StatusNotFound, cardNotFoundError())
		}
	}
}

// GetCards returns all the cards matching the request's requirements
// as JSON
// TODO: Add pagination
func GetCards(c *gin.Context) {

	db := c.MustGet(constants.DB).(*gorm.DB)

	stringSearchParamsToColumnNames := map[string]string{
		"cardtype":       "card_type",
		"clan":           "clan",
		"designillus":    "design_illus",
		"flavor":         "flavor",
		"format":         "format",
		"grade":          "grade",
		"illust":         "illust_",
		"illustcolor":    "illust_color",
		"illust2":        "illust_2",
		"illust3":        "illust_3",
		"illust4":        "illust_4",
		"illust5":        "illust_5",
		"imaginarygift":  "imaginary_gift",
		"italian":        "italian",
		"kana":           "kana",
		"kanji":          "kanji",
		"korean":         "korean",
		"limitationtext": "limitation_text",
		"mangaillust":    "manga_illust",
		"name":           "name",
		"nation":         "nation",
		"note":           "note",
		"othernames":     "other_names",
		"phonetic":       "phonetic",
		"race":           "race",
		"rideskill":      "rideSkill",
		"skill":          "skill",
		"thai":           "thai",
		"translation":    "translation",
		"triggereffect":  "trigger_effect",
	}
	intSearchParamsToColumnNames := map[string]string{
		"critical": "critical",
		"grade":    "grade",
		"power":    "power",
		"shield":   "shield",
	}

	query := db.Preload("Sets").Preload("TournamentStatuses")

	for param, columnName := range stringSearchParamsToColumnNames {
		if val := c.Query(param); val != "" {
			query = query.Where(fmt.Sprintf("UPPER(%s) LIKE ?", columnName), strings.ToUpper(val))
		}
	}

	for param, columnName := range intSearchParamsToColumnNames {
		// TODO: Add support for checking greater than, less than, etc.
		if strVal := c.Query(param); strVal != "" {
			if intVal, err := strconv.Atoi(strVal); err == nil {
				query = query.Where(fmt.Sprintf("%s = ?", columnName), intVal)
			} else {
				c.JSON(http.StatusInternalServerError, invalidIntegerFieldError(param, strVal))
				return
			}
		}
	}

	cards := []models.Card{}
	query.Find(&cards)
	c.JSON(200, cards)
}

// GetCardsInSet returns all the cards in the given set as JSON
func GetCardsInSet(c *gin.Context) {
	// TODO: Move this functionality into the Cards function
	db := c.MustGet(constants.DB).(*gorm.DB)

	set := models.Set{}
	name := c.Query("name")

	if result := db.Preload("Cards.Sets").Preload("Cards.TournamentStatuses").Preload("Cards").
		Model(&models.Set{}).
		Where("UPPER(name) LIKE ?", strings.ToUpper(name)).
		Find(&set); result.RowsAffected > 0 {
		c.JSON(200, set)
	} else {
		c.JSON(http.StatusNotFound, setNotFoundError(name))
	}
}

func invalidCardIDError() map[string]interface{} {
	return gin.H{"error": "please enter a valid card ID"}
}

func cardNotFoundError() map[string]interface{} {
	return gin.H{"error": "no card found"}
}

func invalidIntegerFieldError(param, val string) map[string]interface{} {
	return gin.H{"error": fmt.Sprintf("could not read integer field, %q with value %q", param, val)}
}

func setNotFoundError(name string) map[string]interface{} {
	return gin.H{"error": fmt.Sprintf("no set found with name %q", name)}
}
