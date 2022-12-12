package main

import (
	"go-gin-api/api"
	"go-gin-api/database"
	_ "go-gin-api/docs"
)

func init() {
	database.NewPostGresSQLClient()
}

// @title Swagger Example API
// @version 1.0
func main() {
	r := api.SetupRouter()
	r.Run()
}
