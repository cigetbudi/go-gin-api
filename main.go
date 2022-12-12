package main

import (
	"go-gin-api/api"
	"go-gin-api/database"
	_ "go-gin-api/docs"
)

func init() {
	database.NewPostGresSQLClient()
}

// @title Article API ginSwagger
// @version 1.0
func main() {
	r := api.SetupRouter()
	r.Run()
}
