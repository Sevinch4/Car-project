package postgres

import "database/sql"

type driverRepo struct {
	DB *sql.DB
}

func NewDriverRepo(db *sql.DB) driverRepo {
	return driverRepo{
		DB: db,
	}
}
