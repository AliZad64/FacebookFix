package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "welcome to an amazing api",
	})
}
