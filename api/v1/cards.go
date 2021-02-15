package v1

import (
	"cfv-api/constants"
	"cfv-api/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: Enable querying by URL search params OR JSON
// depending on request header
// TODO: Make queries case insensitive somehow?
func GetCards(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	cardQuery := models.Card{
		CardType:       c.Query("cardtype"),
		Clan:           c.Query("clan"),
		DesignIllus:    c.Query("designillus"),
		Flavor:         c.Query("flavor"),
		Format:         c.Query("format"),
		Grade:          c.Query("grade"),
		Illust:         c.Query("illust"),
		IllustColor:    c.Query("illustcolor"),
		Illust2:        c.Query("illust2"),
		Illust3:        c.Query("illust3"),
		Illust4:        c.Query("illust4"),
		Illust5:        c.Query("illust5"),
		ImaginaryGift:  c.Query("imaginarygift"),
		Italian:        c.Query("italian"),
		Kana:           c.Query("kana"),
		Kanji:          c.Query("kanji"),
		Korean:         c.Query("korean"),
		LimitationText: c.Query("limitationtext"),
		MangaIllust:    c.Query("mangaillust"),
		Name:           c.Query("name"),
		Nation:         c.Query("nation"),
		Note:           c.Query("note"),
		OtherNames:     c.Query("othernames"),
		Phonetic:       c.Query("phonetic"),
		Race:           c.Query("race"),
		RideSkill:      c.Query("rideskill"),
		Skill:          c.Query("skill"),
		Thai:           c.Query("thai"),
		Translation:    c.Query("translation"),
		TriggerEffect:  c.Query("triggereffect"),
	}
	integerFieldNames := []string{"critical", "power", "shield"}
	for _, field := range integerFieldNames {
		if val, err := strconv.Atoi(c.DefaultQuery(field, "-1")); err == nil {
			// Messy but works
			if val != -1 {
				switch field {
				case "critical":
					cardQuery.Critical = val
				case "power":
					cardQuery.Power = val
				case "shield":
					cardQuery.Shield = val
				}
			}
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("could not read integer field, %q with value %q", field, val),
			})
			return
		}
	}
	cards := []models.Card{}
	db.Preload("Sets").Preload("TournamentStatuses").
		Where(&cardQuery).
		Find(&cards)
	c.JSON(200, cards)
}

func GetCardsInSet(c *gin.Context) {
	db := c.MustGet(constants.DB).(*gorm.DB)

	set := models.Set{}
	name := c.Query("name")

	if result := db.Preload("Cards").
		Model(&models.Set{}).
		Where("name ILIKE ?", name).
		Find(&set); result.RowsAffected > 0 {
		c.JSON(200, set)
	} else {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("no set found with name %q", name)})
	}
}
