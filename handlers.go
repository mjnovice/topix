package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func root(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl.html", nil)
}
