package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMemory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "linux",
	})
}
