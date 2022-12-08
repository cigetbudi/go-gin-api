package database

import (
	"fmt"
	"go-gin-api/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error
var a models.Article

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal baca .env ", err)
	}
	return os.Getenv(key)
}

func NewPostGresSQLClient() {
	var (
		host     = getEnv("DB_HOST")
		port     = getEnv("DB_PORT")
		user     = getEnv("DB_USER")
		dbname   = getEnv("DB_NAME")
		password = getEnv("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err = gorm.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(models.Article{})
	fmt.Println("sukses terhubung dengan neondb")
}
