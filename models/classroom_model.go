package models

import "time"

type CreateClassroom struct {
	Name         string   `json:"name"`
	TeacherID    string   `json:"teacherId"`
	AcademicYear string   `json:"academicYear"`
	Subjects     []string `json:"subjects"`
}
type Classroom struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	TeacherID    string    `json:"teacherId"`
	AcademicYear string    `json:"academicYear"`
	CreatedAt    time.Time `json:"createdAt"`
}
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type ClassroomSubject struct {
	ID          string `json:"id"`
	ClassroomID string `json:"classroomId"`
	SubjectID   string `json:"subjectId"`
}
type ClassroomCard struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	TeacherName   string `json:"teacherName"`
	StudentsCount int    `json:"studentsCount"`
	SubjectsCount int    `json:"subjectsCount"`
}
