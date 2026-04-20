package models

import "time"

type TeacherOnboarding struct {
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	
    Email string `json:"email" validate:"required,email"`
    Phone string `json:"phone" validate:"required,e164"`

	Age              int       `json:"age"`
	DateOfBirth      string`json:"dob"`
	Address          string    `json:"address"`
	Qualification    string    `json:"qualification"`
	SubjectsTeaching string    `json:"subjectsTeaching"`
	Password         string    `json:"password"`
}
type UpdateTeacher struct {
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Phone            string    `json:"phone"`
	Age              int       `json:"age"`
	DateOfBirth      time.Time `json:"dob"`
	Address          string    `json:"address"`
	Qualification    string    `json:"qualification"`
	SubjectsTeaching string    `json:"subjectsTeaching"`
}
type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	IsVerified bool   `json:"isVerified"`
	Role       string `json:"role"`
	IsBlocked  bool   `json:"isBlocked"`
}
type CreateSalary struct {
	TeacherID     string `json:"teacherId"`
	BaseSalary    int    `json:"baseSalary"`
	Allowance     int    `json:"allowance"`
	EffectiveFrom string `json:"effectiveFrom"`
}

type SalaryResponse struct {
	TeacherName string `json:"teacherName"`
	BaseSalary  int    `json:"baseSalary"`
	Allowance   int    `json:"allowance"`
	Total       int    `json:"total"`
	Status      string `json:"status"`
}
type UpdateSalary struct {
	BaseSalary    int    `json:"baseSalary"`
	Allowance     int    `json:"allowance"`
	EffectiveFrom string `json:"effectiveFrom"`
}
type TeacherResponse struct {
	ID               string `json:"id"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Age              int    `json:"age"`
	DOB              string `json:"dob"`
	Address          string `json:"address"`
	Qualification    string `json:"qualification"`
	SubjectsTeaching string `json:"subjectsTeaching"`
}
