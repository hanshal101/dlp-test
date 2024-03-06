package controllers

import (
	"dlpTest/database/models"
	"dlpTest/database/postgres"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Name   string
	Rollno int
}

type Out struct {
	Class_Name string
	Class_Room string
	Students   []Result
}

type Out_INP struct {
	Class_Name string
	Class_Room string
	Students   []Result
}

func GET2(c *gin.Context) {
	var syaStudents []models.StudentsDB
	var result []Out

	var alloc []models.AllocationResult
	if err := postgres.DB.Find(&alloc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}

	for _, allocations := range alloc {

		if err := postgres.DB.Where("class = ? AND roll_no BETWEEN ? AND ?", allocations.ClassName, allocations.Start, allocations.End).Find(&syaStudents).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to fetch %s students", allocations.ClassName)})
			return
		}
		var studentNames []Result
		for _, student := range syaStudents {
			studentNames = append(studentNames, Result{student.Name, student.RollNo})
		}

		result = append(result, Out{Class_Name: allocations.ClassName, Class_Room: allocations.ClassRoom, Students: studentNames})

	}

	// Fetch SYA students from the database
	// if err := postgres.DB.Where("class = ? AND roll_no BETWEEN ? AND ?", "SYA", 1, 10).Find(&syaStudents).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch SYA students"})
	// 	return
	// }

	// Create a slice to hold student names
	// var studentNames []Result
	// for _, student := range syaStudents {
	// 	studentNames = append(studentNames, Result{student.Name, student.RollNo})
	// }

	// Create the JSON response
	// response := gin.H{
	// 	"Class":   "SYA",
	// 	"Student": studentNames,
	// }

	c.JSON(http.StatusOK, result)

}
