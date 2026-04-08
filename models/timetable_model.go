package models

type CreateTimetable struct {
	ClassID   string `json:"class_id"`
	Day       string `json:"day"`
	PeriodID  string `json:"period_id"`
	SubjectID string `json:"subject_id"`
	TeacherID string `json:"teacher_id"`
	Type      string `json:"type"` // class / break / lunch
}
