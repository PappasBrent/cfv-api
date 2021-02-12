package main

import (
	"cfv-api/models"
	"fmt"
	"net/http"
	"strconv"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDSN() (string, error) {
	// TODO: Check environment (dev vs. prod) before loading dotenv.
	// Store that as a parameter?
	if err := godotenv.Load(); err != nil {
		return "", err
	}

	host, user, password, db, port := os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB"), os.Getenv("DB_PORT")
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, db, port), nil
}

type CardResult struct {
	*models.Card
	Sets                     []string `json:"Sets"`
	*models.TournamentStatus `json:"TournamentStatus"`
}

func LoadCardResult(card *models.Card, db *gorm.DB) (CardResult, error) {
	result := CardResult{card, []string{}, &models.TournamentStatus{}}
	if err := db.Model(&models.Set{}).
		Select("set.name").
		Where(`set.id IN
				(SELECT set_id
					FROM card_set_xref
					WHERE card_id = ?);`, card.ID).
		Find(&result.Sets).Error; err != nil {
		return result, err
	}
	if err := db.Model(&models.TournamentStatus{}).
		Where("tournament_status.card_id = ?", card.ID).
		Scan(&result.TournamentStatus).Error; err != nil {
		return result, err
	}
	return result, nil
}

func main() {

	dsn, err := LoadDSN()

	if err != nil {
		fmt.Println(err)
		return
	}

	dialector := postgres.New(postgres.Config{DSN: dsn, PreferSimpleProtocol: true})
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		return
	}

	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		if id, err := strconv.ParseInt(c.Query("id"), 10, 64); err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "invalid query parameter"})
		} else {
			card := models.Card{}
			if err := db.Model(&models.Card{}).Where("id = ?", id).First(&card).Error; err != nil {
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "card not found"})
				return
			}
			if result, err := LoadCardResult(&card, db); err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "error obtaining associated sets and tournaments statuses for card"})
				return
			} else {
				c.JSON(200, result)
			}
		}
	})

	app.Run()
}
