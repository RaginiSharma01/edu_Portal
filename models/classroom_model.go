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

type AddStudentsToClassroom struct {
	StudentIDs []string `json:"studentIds"`
}

type ClassroomDetail struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	TeacherName  string           `json:"teacherName"`
	AcademicYear string           `json:"academicYear"`
	Subjects     []string         `json:"subjects"`
	Students     []StudentInClass `json:"students"`
}

type StudentInClass struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
