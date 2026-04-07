package db

import (
	"smp/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ProvidePgDb(cfg *config.Config) *PgDb {
	db, err := ConnectDb(cfg)
	if err != nil {
		panic(err)
	}
	return db
}

func ProvidePool(db *PgDb) *pgxpool.Pool {
	return db.Pool
}
