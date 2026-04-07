package repository

import (
	"context"
	"smp/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EventRepo struct {
	DB *pgxpool.Pool
}

func NewEventRepository(pool *pgxpool.Pool) *EventRepo {
	return &EventRepo{
		DB: pool,
	}
}

func (r *EventRepo) CreateEvent(ctx context.Context, req models.CreateEvent) (string, error) {

	query := `
	INSERT INTO events (title, event_date, venue, description, type)
	VALUES ($1,$2,$3,$4,$5)
	RETURNING id
	`

	var eventID string

	err := r.DB.QueryRow(ctx, query,
		req.Title,
		req.EventDate,
		req.Venue,
		req.Description,
		req.Type,
	).Scan(&eventID)

	if err != nil {
		return "", err
	}

	return eventID, nil
}

func (r *EventRepo) GetEvents(ctx context.Context) ([]models.Event, error) {

	query := `
	SELECT id, title, event_date, venue, description, type
	FROM events
	ORDER BY event_date
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var e models.Event

		err := rows.Scan(
			&e.ID,
			&e.Title,
			&e.EventDate,
			&e.Venue,
			&e.Description,
			&e.Type,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return events, nil
}
