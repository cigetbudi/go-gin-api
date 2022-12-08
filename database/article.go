package database

import (
	"errors"
	"go-gin-api/models"
)

func CreateArticle(a *models.Article) (*models.Article, error) {
	res := db.Create(a)
	if res.RowsAffected == 0 {
		return &models.Article{}, errors.New("article gagal ditambahkan")
	}
	return a, nil
}

func ReadArticle(id string) (*models.Article, error) {
	var a models.Article
	res := db.First(&a, id)
	if res.RowsAffected == 0 {
		return nil, errors.New("article tidak ditemukan")
	}
	return &a, nil
}

func ReadArticles() ([]*models.Article, error) {
	var as []*models.Article
	res := db.Find(&as)
	if res.Error != nil {
		return nil, errors.New("article tidak ditemukan")
	}
	return as, nil
}

func UpdateArticle(a *models.Article) (*models.Article, error) {
	var upa models.Article
	res := db.Model(&upa).Where(a.ID).Updates(a)
	if res.RowsAffected == 0 {
		return &models.Article{}, errors.New("article gagal diupdate")
	}
	return &upa, nil
}

func DeleteArticle(id string) error {
	var dela models.Article
	res := db.Where(id).Delete(&dela)
	if res.RowsAffected == 0 {
		return errors.New("article gagal dihapus")
	}
	return nil
}
