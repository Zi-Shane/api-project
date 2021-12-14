package controllers

import (
	"api/database"
	"api/services"

	"github.com/gin-gonic/gin"
)

// HandlerFunc of `GET /api/language/{id}`
func GetLanguage(c *gin.Context) {
	queryId := c.Param("id") // 獲取url參數

	// Call function from `services` to get data
	var languages []database.Languages = services.ReadLanguage(queryId)

	// Response data to Client
	c.JSON(200, gin.H{
		"result": languages,
	})
}
