package models

import "time"

type DashboardSummary struct {
	Students     int `json:"students"`
	Subjects     int `json:"subjects"`
	PendingMarks int `json:"pendingMarks"`
	Events       int `json:"events"`
}

type DashboardEvent struct {
	Title    string    `json:"title"`
	Date     time.Time `json:"date"`
	Location string    `json:"location"`
}

type Activity struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

type TeacherDashboard struct {
	Summary          DashboardSummary `json:"summary"`
	UpcomingEvents   []DashboardEvent `json:"upcomingEvents"`
	RecentActivities []Activity       `json:"recentActivities"`
}

type DashboardRequest struct {
	TeacherID string `json:"teacherId"`
}
