package models

type SubjectMarks struct {
	SubjectID string `json:"subjectId"`
	Marks     int    `json:"marks"`
}

type CreateMarks struct {
	StudentID string         `json:"studentId"`
	Term      string         `json:"term"`
	Marks     []SubjectMarks `json:"marks"`
}

type StudentMarks struct {
	Student     string  `json:"student"`
	Mathematics *int    `json:"mathematics"`
	Physics     *int    `json:"physics"`
	Chemistry   *int    `json:"chemistry"`
	English     *int    `json:"english"`
	Hindi       *int    `json:"hindi"`
	Social      *int    `json:"social"`
	Total       int     `json:"total"`
	Percentage  float64 `json:"percentage"`
}
