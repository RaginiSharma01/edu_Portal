package repository

import (
	"context"
	"strings"

	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherDashboardRepository struct {
	DB *pgxpool.Pool
}

func NewTeacherDashboardRepository(pool *pgxpool.Pool) *TeacherDashboardRepository {
	return &TeacherDashboardRepository{DB: pool}
}

func (r *TeacherDashboardRepository) GetDashboard(ctx context.Context, teacherID string) (*models.TeacherDashboard, error) {
	var dashboard models.TeacherDashboard

	//Students 
	err := r.DB.QueryRow(ctx, `
		SELECT COUNT(s.user_id)
		FROM students s
		JOIN classrooms c ON s.classroom_id = c.id
		WHERE c.teacher_id = $1
	`, teacherID).Scan(&dashboard.Summary.Students)
	if err != nil {
		return nil, err
	}

	// Subjects 
	var subjects string
	err = r.DB.QueryRow(ctx, `
		SELECT subjects_teaching FROM teachers WHERE user_id = $1
	`, teacherID).Scan(&subjects)
	if err == nil && subjects != "" {
		dashboard.Summary.Subjects = len(strings.Split(subjects, ","))
	}

	//  Events Count 
	err = r.DB.QueryRow(ctx, `
		SELECT COUNT(*) FROM events
		WHERE event_date >= CURRENT_DATE
		AND event_date <= CURRENT_DATE + INTERVAL '7 days'
	`).Scan(&dashboard.Summary.Events)
	if err != nil {
		return nil, err
	}

	//  Upcoming Events
	rows, err := r.DB.Query(ctx, `
		SELECT title, event_date, venue
		FROM events
		WHERE event_date >= CURRENT_DATE
		ORDER BY event_date ASC
		LIMIT 5
	`)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			var e models.DashboardEvent
			rows.Scan(&e.Title, &e.Date, &e.Location)
			dashboard.UpcomingEvents = append(dashboard.UpcomingEvents, e)
		}
	}

	//Recent Activities 
	rows, err = r.DB.Query(ctx, `
		SELECT type, message
		FROM activities
		ORDER BY created_at DESC
		LIMIT 5
	`)
	if err == nil {
		defer rows.Close()

		for rows.Next() {
			var a models.Activity
			rows.Scan(&a.Type, &a.Message)
			dashboard.RecentActivities = append(dashboard.RecentActivities, a)
		}
	}

	//  Pending Marks 
	dashboard.Summary.PendingMarks = 0

	return &dashboard, nil
}