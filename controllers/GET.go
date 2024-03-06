package controllers

import (
	"dlpTest/database/models"
	"dlpTest/database/postgres"

	"github.com/gin-gonic/gin"
)

func GetGo(c *gin.Context) {
	var result models.AllocationResult
	if err := postgres.DB.Where(&result, "").Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, result)
}
