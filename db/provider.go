package db

import (
	"log"
	"smp/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ProvidePgDb(cfg *config.Config) *PgDb {
	db, err := ConnectDb(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func ProvidePool(db *PgDb) *pgxpool.Pool {
	return db.Pool
}
