package main

import (
	"go-gin-api/api"
	"go-gin-api/database"
)

func init() {
	database.NewPostGresSQLClient()
}

func main() {
	r := api.SetupRouter()
	r.Run(":6821")
}
