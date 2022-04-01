package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/garcia-paulo/go-rest/application/servicers"
)

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(servicers.GetQuotes())
}

func GetQuotesById(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(servicers.GetQuotesById(r))
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(servicers.CreateQuote(r))
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	servicers.DeleteQuote(r)
}

func UpdateQuote(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(servicers.UpdateQuote(r))
}
