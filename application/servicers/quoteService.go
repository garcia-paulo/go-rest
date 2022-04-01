package servicers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/garcia-paulo/go-rest/domain/models"
	"github.com/garcia-paulo/go-rest/infra/repositories"
	"github.com/gorilla/mux"
)

func GetQuotes() []models.Quote {
	quotes := repositories.GetQuotes()
	return quotes
}

func GetQuotesById(r *http.Request) models.Quote {
	vars := mux.Vars(r)
	quoteId, err := strconv.Atoi(vars["quoteId"])
	if err != nil {
		panic(err.Error())
	}

	quote := repositories.GetQuotesById(quoteId)
	return quote
}

func CreateQuote(r *http.Request) models.Quote {
	quote := models.Quote{}
	json.NewDecoder(r.Body).Decode(&quote)
	repositories.CreateQuote(&quote)
	return quote
}

func DeleteQuote(r *http.Request) {
	vars := mux.Vars(r)
	quoteId, err := strconv.Atoi(vars["quoteId"])
	if err != nil {
		panic(err.Error())
	}

	repositories.DeleteQuote(quoteId)
}

func UpdateQuote(r *http.Request) models.Quote {
	vars := mux.Vars(r)
	quoteId, err := strconv.Atoi(vars["quoteId"])
	if err != nil {
		panic(err.Error())
	}

	quote := models.Quote{}
	json.NewDecoder(r.Body).Decode(&quote)
	repositories.UpdateQuote(&quote, quoteId)

	return quote
}
