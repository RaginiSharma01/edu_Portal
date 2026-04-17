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

// type StudentMarks struct {
// 	Student    string
// 	Math       int32
// 	Science    int32
// 	Hindi      int32
// 	English    int32
// 	Computer   int32
// 	Social     int32
// 	Total      int64
// 	Percentage float64
// }

type StudentMarks struct {
	Student    string         `json:"student"`
	Subjects   map[string]int `json:"subjects"`
	Total      int            `json:"total"`
	MaxTotal   int            `json:"max_total"`
	Percentage float64        `json:"percentage"`
}
