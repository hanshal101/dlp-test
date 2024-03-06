package controllers

import (
	"dlpTest/database/models"
	"dlpTest/database/postgres"
	"dlpTest/helpers"
	"fmt"
	"log"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DualAlloc(c *gin.Context) {
	var reqArr models.ReqArr
	if err := c.BindJSON(&reqArr); err != nil {
		log.Fatalf("Error in Binding")
	}

	result := []models.AllocationResult{}

	for _, req := range reqArr.ReqAll {
		totalCapacity := 0
		for _, cap := range req.Class {
			totalCapacity += int(cap.Capacity)
		}

		max_no := math.Max(float64(req.Class1.Strength), float64(req.Class2.Strength))
		if max_no > float64(totalCapacity) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient Space"})
		}

		roomRollNo := 1
		tx := postgres.DB.Begin()

		for i, class := range req.Class {
			result = append(result, models.AllocationResult{
				ClassRoom: class.Room,
				ClassName: req.Class1.Name,
				Start:     int64(roomRollNo),
				End:       helpers.Allot(req.Class1.Strength, class.Capacity, i),
			})
			result = append(result, models.AllocationResult{
				ClassRoom: class.Room,
				ClassName: req.Class2.Name,
				Start:     int64(roomRollNo),
				End:       helpers.Allot(req.Class2.Strength, class.Capacity, i),
			})

			roomRollNo += int(class.Capacity)
		}

		if err := tx.Create(&result).Error; err != nil {
			tx.Rollback()
			fmt.Printf("Error : %v\n", err)
			return
		}

		tx.Commit()
	}

	c.JSON(http.StatusOK, result)
}
