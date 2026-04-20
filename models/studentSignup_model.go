package models

import "time"

type StudentOnboarding struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Age         int    `json:"age"`
	DateOfBirth string `json:"dob"`
	Address     string `json:"address"`

	FatherName   string `json:"fatherName"`
	MotherName   string `json:"motherName"`
	GuardianName string `json:"guardianName"`
	Occupation   string `json:"occupation"`

	Height int `json:"height"`
	Weight int `json:"weight"`

	Password string `json:"password"`
}
type UpdateStudent struct {
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Phone        string    `json:"phone"`
	Age          int       `json:"age"`
	DateOfBirth  time.Time `json:"dob"`
	Address      string    `json:"address"`
	FatherName   string    `json:"fatherName"`
	MotherName   string    `json:"motherName"`
	GuardianName string    `json:"guardianName"`
	Occupation   string    `json:"occupation"`
	Height       int       `json:"height"`
	Weight       int       `json:"weight"`
}

type Student struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Age          int       `json:"age"`
	DateOfBirth  time.Time `json:"dob"`
	Address      string    `json:"address"`
	FatherName   string    `json:"fatherName"`
	MotherName   string    `json:"motherName"`
	GuardianName string    `json:"guardianName"`
	Occupation   string    `json:"occupation"`
	Height       int       `json:"height"`
	Weight       int       `json:"weight"`
}
