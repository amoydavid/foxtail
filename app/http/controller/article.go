package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleList(c *gin.Context) {
	b := "hello posts here"
	c.String(http.StatusOK, "%s", b)
}
