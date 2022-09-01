package product

import (
	"database/sql"
	"fmt"
	"os"
)

type Repository interface {
	Hello() string
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Hello() string {
	var greeting string
	err := r.db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return greeting
}
