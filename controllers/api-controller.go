package controllers

import (
	"api/database"
	"api/services"
	"os"

	"github.com/gin-gonic/gin"
)

func GetHostname(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.JSON(200, gin.H{
		"hostname": hostname,
	})
}

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

func AddLanguage(c *gin.Context) {
	var m []database.Languages
	c.Bind(&m)
	err, rowsAffected := services.InsertLanguage(m)
	if err != nil {
		c.JSON(200, gin.H{"message": err.Error()})
	} else {
		c.JSON(200, gin.H{"message": "success", "rowsAffected": rowsAffected})
	}
}

func RemoveLanguage(c *gin.Context) {
	queryId := c.Param("id")
	err, rowsAffected := services.DeleteLanguage(queryId)
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

// func UpdateLanguage(c *gin.Context) {

// 	c.JSON(200, "success")
// }
