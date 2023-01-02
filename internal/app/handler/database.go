package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	pgxpool *pgxpool.Pool
}

func NewDatabase(pgxpool *pgxpool.Pool) *Database {
	return &Database{pgxpool: pgxpool}
}

func (db *Database) PingPostreSQL(w http.ResponseWriter, r *http.Request) {
	if err := db.pgxpool.Ping(r.Context()); err != nil {
		http.Error(w, errServerInternal.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
