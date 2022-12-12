package api

import (
	"go-gin-api/database"
	"go-gin-api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostArticle             godoc
// @Summary      Post a new Article
// @Description  Takes a article JSON and store in DB. Return saved JSON.
// @Tags         article
// @Produce      json
// @Param        article  body      models.Article  true  "Article JSON"
// @Success      200   {object}  models.Article
// @Router       /api/v1/articles [post]
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

// GetArticle           godoc
// @Summary      Get Article By id
// @Description  Responds with the list of all article as JSON.
// @Tags         article
// @Produce      json
// @Param        id     path    string  true  "ID"
// @Success      200    json    models.Article
// @Failure      404    json    gin.H
// @Router       /api/v1/articles/{id} [get]
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

// GetArticles           godoc
// @Summary      Get Article array
// @Description  Responds with the list of all articles as JSON.
// @Tags         article
// @Produce      json
// @Success      200  {array}  []models.Article
// @Router       /api/v1/articles [get]
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

// PutArticle             godoc
// @Summary      Update an Article
// @Description  Takes a article JSON and store in DB. Return saved JSON.
// @Tags         article
// @Produce      json
// @Param        id     path    string  true  "ID"
// @Param        article  body      models.Article  true  "Article JSON"
// @Success      200   {object}  models.Article
// @Router       /api/v1/articles/{id} [put]
func putArticle(c *gin.Context) {
	var a models.Article
	id := c.Param("id")
	idInt, _ := strconv.ParseInt(id, 10, 64)
	a.ID = uint(idInt)

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

// DeleteArticle           godoc
// @Summary      Delete Article By id
// @Description  Responds with JSON message.
// @Tags         article
// @Produce      json
// @Param        id     path    string  true  "ID"
// @Success      200    json    gin.H
// @Failure      404    json    gin.H
// @Router       /api/v1/articles/{id} [delete]
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
