package repository

//student dashboard

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentDashboardRepository struct {
	DB *pgxpool.Pool
}

func NewStudentDashboardRepository(pool *pgxpool.Pool) *StudentDashboardRepository {
	return &StudentDashboardRepository{DB: pool}
}

func (r *StudentDashboardRepository) GetDashboard(ctx context.Context, studentID string) (*models.StudentDashboard, error) {

	var dashboard models.StudentDashboard

	// Subjects count (based on student's class)
	err := r.DB.QueryRow(ctx, `
		SELECT COUNT(cs.subject_id)
		FROM classroom_subjects cs
		JOIN students s ON s.classroom_id = cs.classroom_id
		WHERE s.user_id = $1
	`, studentID).Scan(&dashboard.Summary.Subjects)
	if err != nil {
		return nil, err
	}

	// Events this week
	err = r.DB.QueryRow(ctx, `
		SELECT COUNT(*)
		FROM events
		WHERE event_date >= CURRENT_DATE
		AND event_date <= CURRENT_DATE + INTERVAL '7 days'
	`).Scan(&dashboard.Summary.Events)
	if err != nil {
		return nil, err
	}

	// Upcoming Events
	rows, err := r.DB.Query(ctx, `
		SELECT title, event_date, venue
		FROM events
		WHERE event_date >= CURRENT_DATE
		ORDER BY event_date ASC
		LIMIT 5
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var e models.DashboardEvent
		rows.Scan(&e.Title, &e.Date, &e.Location)
		dashboard.UpcomingEvents = append(dashboard.UpcomingEvents, e)
	}

	return &dashboard, nil
}
