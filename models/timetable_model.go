package models

type CreateTimetable struct {
	ClassID   string  `json:"classId"`
	Day       string  `json:"day"`
	PeriodID  string  `json:"periodId"`
	SubjectID *string `json:"subjectId"`
	TeacherID *string `json:"teacherId"`
	Type      string  `json:"type"`
}

type TimetableRow struct {
	PeriodNumber int    `json:"periodNumber"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`

	Monday    *string `json:"monday"`
	Tuesday   *string `json:"tuesday"`
	Wednesday *string `json:"wednesday"`
	Thursday  *string `json:"thursday"`
	Friday    *string `json:"friday"`
}
