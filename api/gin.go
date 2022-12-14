package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
