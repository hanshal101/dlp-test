package models

import "gorm.io/gorm"

type Class struct {
	Room     string `json:"room"`
	Capacity int64  `json:"capacity"`
}

type Req struct {
	gorm.Model
	Class1 Year    `json:"class1"`
	Class2 Year    `json:"class2"`
	Class  []Class `json:"class"`
}

type ReqArr struct {
	ReqAll []Req `json:"reqAll"`
}

type Year struct {
	Name     string `json:"name"`
	Strength int64  `json:"strength"`
}

type AllocationResult struct {
	gorm.Model
	ClassRoom string `json:"classroom"`
	ClassName string `json:"classname"`
	Start     int64  `json:"start"`
	End       int64  `json:"end"`
}

type Students struct {
	Name       string `json:"name"`
	RollNo     int    `json:"roll_no"`
	Email      string `json:"email"`
	Class      string `json:"class"`
	Department string `json:"department"`
}

type StudentsDB struct {
	gorm.Model
	Name       string `json:"name"`
	RollNo     int    `json:"roll_no"`
	Email      string `json:"email"`
	Class      string `json:"class"`
	Department string `json:"department"`
}
