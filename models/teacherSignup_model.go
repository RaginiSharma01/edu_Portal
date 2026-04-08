package models

type TeacherOnboarding struct {
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	Age              int    `json:"age"`
	DateOfBirth      string `json:"dob"`
	Address          string `json:"address"`
	Qualification    string `json:"qualification"`
	SubjectsTeaching string `json:"subjectsTeaching"`
	Password         string `json:"password"`
}

type User struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Password   string `json:"-"`
	IsVerified bool   `json:"isVerified"`
	Role       string `json:"role"`
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
