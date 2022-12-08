package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

type Article struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        int    `json:"rate"`
}

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

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
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

	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article gagal ditambahkan")
	}
	return a, nil
}
