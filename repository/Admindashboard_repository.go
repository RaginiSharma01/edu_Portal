// dashboard of Admin
// dashboard of students
// dasdboard of teachers
package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AdminDashboardRepository struct {
	DB *pgxpool.Pool
}

func NewDashboardRepository(pool *pgxpool.Pool) *AdminDashboardRepository {
	return &AdminDashboardRepository{DB: pool}
}

// SUMMARY 
func (r *AdminDashboardRepository) GetSummary() (map[string]int, error) {
	var students, teachers, classrooms, events int

	err := r.DB.QueryRow(context.Background(), `
		SELECT 
			(SELECT COUNT(*) FROM students),
			(SELECT COUNT(*) FROM teachers),
			(SELECT COUNT(*) FROM classrooms),
			(SELECT COUNT(*) FROM events WHERE event_date >= CURRENT_DATE)
	`).Scan(&students, &teachers, &classrooms, &events)

	if err != nil {
		return nil, err
	}

	return map[string]int{
		"students":   students,
		"teachers":   teachers,
		"classrooms": classrooms,
		"events":     events,
	}, nil
}

// ================= RECENT ACTIVITY =================
func (r *AdminDashboardRepository) GetRecentActivities() ([]string, error) {

	rows, err := r.DB.Query(context.Background(), `
		SELECT message
		FROM activities
		ORDER BY created_at DESC
		LIMIT 5
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []string

	for rows.Next() {
		var msg string
		if err := rows.Scan(&msg); err != nil {
			return nil, err
		}
		activities = append(activities, msg)
	}

	return activities, nil
}

// ================= EVENTS =================
func (r *AdminDashboardRepository) GetUpcomingEvents() ([]models.Events, error) {

	rows, err := r.DB.Query(context.Background(), `
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

	var events []models.Events

	for rows.Next() {
		var e models.Events
		if err := rows.Scan(&e.Title, &e.Date, &e.Location); err != nil {
			return nil, err
		}
		events = append(events, e)
	}

	return events, nil
}

// ================= FINAL DASHBOARD =================
func (r *AdminDashboardRepository) GetDashboard() (map[string]interface{}, error) {

	summary, err := r.GetSummary()
	if err != nil {
		return nil, err
	}

	activities, err := r.GetRecentActivities()
	if err != nil {
		return nil, err
	}

	events, err := r.GetUpcomingEvents()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"summary":          summary,
		"recentActivities": activities,
		"upcomingEvents":   events,
	}, nil
}
