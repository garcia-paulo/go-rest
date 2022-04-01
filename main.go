package main

import (
	"github.com/garcia-paulo/go-rest/infra/database"
	"github.com/garcia-paulo/go-rest/presentation/routes"
)

func main() {
	database.DBConnect()
	routes.HandleRequest()
}
