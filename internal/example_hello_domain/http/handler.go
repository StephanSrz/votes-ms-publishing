package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(202, gin.H{"msg": "Hello"})
}

func HelloName(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	requestBody := string(body)
	c.JSON(202, gin.H{"msg": requestBody})
}
