package models

import "time"

type CreateClassroom struct {
	Name         string   `json:"name"`
	TeacherID    string   `json:"teacher_id"`
	AcademicYear string   `json:"academic_year"`
	Subjects     []string `json:"subjects"`
}
type Classroom struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TeacherID    string    `json:"teacher_id"`
	AcademicYear string    `json:"academic_year"`
	CreatedAt    time.Time `json:"created_at"`
}
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ClassroomSubject struct {
	ID          string `json:"id"`
	ClassroomID string `json:"classroom_id"`
	SubjectID   string `json:"subject_id"`
}
