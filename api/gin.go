package api

import (
	"go-gin-api/database"
	"go-gin-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", home)
	r.GET("/api/v1/articles/:id", getArticle)
	r.GET("/api/v1/articles", getArticles)
	r.POST("/api/v1/articles", postArticle)
	r.PUT("/api/v1/articles/:id", putArticle)
	r.DELETE("/api/v1/articles/:id", deleteArticle)
	return r
}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Selamat datang di API gin-gorm article",
	})
}

func postArticle(c *gin.Context) {
	var a models.Article
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.CreateArticle(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"article": res,
	})

}

func getArticle(c *gin.Context) {
	id := c.Param("id")

	a, err := database.ReadArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article tidak ditemukan",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"article": a,
	})
}

func getArticles(c *gin.Context) {
	as, err := database.ReadArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"articles": as,
	})
}

func putArticle(c *gin.Context) {
	var a models.Article
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	res, err := database.UpdateArticle(&a)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"article": res,
	})
}

func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "article tidak ditemukan, gagal dihapus",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "article berhasil dihapus",
	})

}
