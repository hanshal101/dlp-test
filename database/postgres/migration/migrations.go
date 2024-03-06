package main

import (
	"dlpTest/database/models"
	"dlpTest/database/postgres"
)

func init() {
	postgres.PostgresInitializer()
}

func main() {
	postgres.DB.AutoMigrate(&models.AllocationResult{})
	postgres.DB.AutoMigrate(&models.StudentsDB{})
}
