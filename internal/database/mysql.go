package database

import (
	"LearnECHO/internal/config"
	"database/sql"
	"fmt"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp/%s",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)
	db, err := sql.Open(cfg.DBDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}
