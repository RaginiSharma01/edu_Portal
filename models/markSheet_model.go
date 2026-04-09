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
	Student    string
	Math       *int
	Science    *int
	Hindi      *int
	English    *int
	Computer   *int
	Social     *int
	Total      int
	Percentage float64
}
