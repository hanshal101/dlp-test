package main

import (
	"dlpTest/database/models"
	"dlpTest/database/postgres"
	"fmt"
)

func init() {
	postgres.PostgresInitializer()
}

func main() {
	data := []models.Students{
		{Name: "Alice", RollNo: 1, Email: "alice@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Bob", RollNo: 2, Email: "bob@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Charlie", RollNo: 3, Email: "charlie@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "David", RollNo: 4, Email: "david@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Eve", RollNo: 5, Email: "eve@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Frank", RollNo: 6, Email: "frank@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Grace", RollNo: 7, Email: "grace@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Henry", RollNo: 8, Email: "henry@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Ivy", RollNo: 9, Email: "ivy@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Jack", RollNo: 10, Email: "jack@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Kate", RollNo: 11, Email: "kate@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Leo", RollNo: 12, Email: "leo@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Mia", RollNo: 13, Email: "mia@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Nick", RollNo: 14, Email: "nick@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Olivia", RollNo: 15, Email: "olivia@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Peter", RollNo: 16, Email: "peter@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Queen", RollNo: 17, Email: "queen@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Ryan", RollNo: 18, Email: "ryan@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Sara", RollNo: 19, Email: "sara@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Tom", RollNo: 20, Email: "tom@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Uma", RollNo: 21, Email: "uma@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Victor", RollNo: 22, Email: "victor@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Wendy", RollNo: 23, Email: "wendy@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Xavier", RollNo: 24, Email: "xavier@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Yara", RollNo: 25, Email: "yara@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Zane", RollNo: 26, Email: "zane@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Adam", RollNo: 27, Email: "adam@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Ben", RollNo: 28, Email: "ben@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Cathy", RollNo: 29, Email: "cathy@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Dylan", RollNo: 30, Email: "dylan@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Emily", RollNo: 31, Email: "emily@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Finn", RollNo: 32, Email: "finn@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Gina", RollNo: 33, Email: "gina@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Hank", RollNo: 34, Email: "hank@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Ian", RollNo: 35, Email: "ian@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Jill", RollNo: 36, Email: "jill@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Kurt", RollNo: 37, Email: "kurt@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Lara", RollNo: 38, Email: "lara@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Mike", RollNo: 39, Email: "mike@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Nora", RollNo: 40, Email: "nora@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Omar", RollNo: 41, Email: "omar@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Patty", RollNo: 42, Email: "patty@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Quinn", RollNo: 43, Email: "quinn@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Ravi", RollNo: 44, Email: "ravi@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Sandy", RollNo: 45, Email: "sandy@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Tina", RollNo: 46, Email: "tina@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Umar", RollNo: 47, Email: "umar@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Vera", RollNo: 48, Email: "vera@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Walt", RollNo: 49, Email: "walt@example.com", Class: "SYA", Department: "EXTC"},
		{Name: "Xena", RollNo: 50, Email: "xena@example.com", Class: "SYA", Department: "EXTC"},
	}

	for _, student := range data {
		studentDB := models.StudentsDB{
			Name:       student.Name,
			RollNo:     student.RollNo,
			Email:      student.Email,
			Class:      student.Class,
			Department: student.Department,
		}
		result := postgres.DB.Create(&studentDB)
		if result.Error != nil {
			fmt.Println("Error:", result.Error)
		}
	}
	fmt.Println("Data inserted successfully")

}
