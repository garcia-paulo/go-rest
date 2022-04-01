package repositories

import (
	"github.com/garcia-paulo/go-rest/domain/models"
	"github.com/garcia-paulo/go-rest/infra/database"
)

func GetQuotes() []models.Quote {
	quotes := []models.Quote{}
	database.DB.Find(&quotes)
	return quotes
}

func GetQuotesById(quoteId int) models.Quote {
	quote := models.Quote{}
	database.DB.First(&quote, quoteId)
	return quote
}

func CreateQuote(quote *models.Quote) {
	database.DB.Create(&quote)
}

func DeleteQuote(quoteId int) {
	database.DB.Delete(&models.Quote{}, quoteId)
}

func UpdateQuote(quote *models.Quote, quoteId int) {
	quote.Id = quoteId
	database.DB.Save(&quote)
}
