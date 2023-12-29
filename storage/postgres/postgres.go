package postgres

import (
	"carProject/config"
	"database/sql"
	"fmt"
)

type Store struct {
	DB         *sql.DB
	CarStorage carRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host = %s port = %s user = %s password = %s database = %s sslmode=disable`,
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	carR := NewCarRepo(db)
	return Store{
		DB:         db,
		CarStorage: carR,
	}, nil
}
