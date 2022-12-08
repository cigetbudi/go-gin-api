package api

import (
	"go-gin-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Selamat datang di API gin-gorm article",
	})
}

func postArticle(c *gin.Context) {
	var a database.Article
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
