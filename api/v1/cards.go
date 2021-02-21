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

// TODO: Enable querying by URL search params OR JSON
// depending on request header

// GetCard returns a single card as JSON
func GetCard(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	if id, err := strconv.Atoi(c.Query("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "please enter a valid card ID"})
	} else {
		cardResult := models.Card{}

		if result := db.Preload("Sets").Preload("TournamentStatuses").
			Where("id = ?", id).
			Find(&cardResult); result.RowsAffected == 1 {
			c.JSON(200, cardResult)
		} else {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "no card found"})
		}
	}
}

// GetCards returns all the cards matching the request's requirements
// as JSON
func GetCards(c *gin.Context) {

	if id := c.Query("id"); id != "" {
		GetCard(c)
		return
	}

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
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": fmt.Sprintf("could not read integer field, %q with value %q", param, strVal),
				})
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
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no set found with name %q", name)})
	}
}
