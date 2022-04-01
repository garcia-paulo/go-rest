package routes

import (
	"log"
	"net/http"

	"github.com/garcia-paulo/go-rest/application/middlewares"
	"github.com/garcia-paulo/go-rest/presentation/controllers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.SetContentType)

	r.HandleFunc("/quotes", controllers.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes", controllers.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes/{quoteId}", controllers.GetQuotesById).Methods("GET")
	r.HandleFunc("/quotes/{quoteId}", controllers.DeleteQuote).Methods("DELETE")
	r.HandleFunc("/quotes/{quoteId}", controllers.UpdateQuote).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
