package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GetHostname(c *gin.Context) {
	hostname, _ := os.Hostname()
	c.JSON(200, gin.H{
		"hostname": hostname,
	})
}
