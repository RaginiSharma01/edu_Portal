package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func LogActivity(db *pgxpool.Pool, message string, activityType string) {
	if db == nil { 
		fmt.Println("LogActivity: db is nil, skipping")
		return
	}
	_, err := db.Exec(context.Background(), `
        INSERT INTO activities (type, message) 
        VALUES ($1, $2)
    `, activityType, message)

	if err != nil {
		fmt.Println("activity log failed:", err)
	}
}
