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
	language, err := services.ReadLanguage(queryId)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		// Response data to Client
		c.JSON(200, gin.H{"message": "success", "result": language})
	}
}

func GetLanguages(c *gin.Context) {
	startId := c.Param("start")
	endId := c.Param("end")
	// Call function from `services` to get data
	languages, err := services.ReadLanguages(startId, endId)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		// Response data to Client
		c.JSON(200, gin.H{"message": "success", "result": languages})
	}
}

func AddLanguage(c *gin.Context) {
	var m []database.Languages
	c.Bind(&m)
	rowsAffected, err := services.InsertLanguage(m)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success", "rowsAffected": rowsAffected})
	}
}

func RemoveLanguage(c *gin.Context) {
	queryId := c.Param("language")
	rowsAffected, err := services.DeleteLanguage(queryId)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		var mesg string
		if rowsAffected == 0 {
			mesg = "Language not found"
		} else {
			mesg = "success"
		}
		c.JSON(200, gin.H{"message": mesg, "rowsAffected": rowsAffected})
	}
}

func UpdateLanguage(c *gin.Context) {
	var m database.Languages
	c.Bind(&m)
	rowsAffected, err := services.UpdateLanguage(m)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success", "rowsAffected": rowsAffected})
	}
}
