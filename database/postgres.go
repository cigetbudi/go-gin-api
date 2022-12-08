package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	db.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article gagal ditambahkan")
	}
	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var a Article
	res := db.First(&a, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("article tidak ditemukan")
	}
	return &a, nil
}

func ReadArticles() ([]*Article, error) {
	var as []*Article
	res := db.Find(&as)
	if res.Error != nil {
		return nil, errors.New("article tidak ditemukan")
	}
	return as, nil
}

func UpdateArticle(a *Article) (*Article, error) {
	var upa Article
	res := db.Model(&upa).Where(a.ID).Updates(a)
	if res.RowsAffected == 0 {
		return &Article{}, errors.New("article gagal diupdate")
	}
	return &upa, nil
}

func DeleteArticle(id string) error {
	var dela Article
	res := db.Where(id).Delete(&dela)
	if res.RowsAffected == 0 {
		return errors.New("article gagal dihapus")
	}
	return nil
}
