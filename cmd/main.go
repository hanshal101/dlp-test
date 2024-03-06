package main

import (
	"dlpTest/controllers"
	"dlpTest/database/postgres"

	"github.com/gin-gonic/gin"
)

// type Students struct {
// 	Name       string
// 	RollNo     int
// 	Email      string
// 	Class      string
// 	Department string
// }

func init() {
	postgres.PostgresInitializer()
}

func main() {

	r := gin.Default()

	r.POST("/dualAllocation", controllers.DualAlloc)
	r.POST("/singleAllocation")
	r.GET("/", controllers.GetGo)
	r.GET("/st", controllers.GET2)
	r.Run(":9876")
}
