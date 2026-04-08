package models

type CreateTimetable struct {
	ClassID   string `json:"classId"`
	Day       string `json:"day"`
	PeriodID  string `json:"periodId"`
	SubjectID string `json:"subjectId"`
	TeacherID string `json:"teacherId"`
	Type      string `json:"type"` // class / break / lunch
}

type TimetableRow struct {
	PeriodNumber int    `json:"period_number"`
	StartTime    string `json:"start_time"`
	EndTime      string `json:"end_time"`

	Monday    *string `json:"monday"`
	Tuesday   *string `json:"tuesday"`
	Wednesday *string `json:"wednesday"`
	Thursday  *string `json:"thursday"`
	Friday    *string `json:"friday"`
}
