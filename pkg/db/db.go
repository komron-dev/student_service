package db

import (
	"fmt"

	"student_service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //postgres drivers
)

func ConnectToDbAndAlsoForSuite(cfg config.Config) (*sqlx.DB, error, func()) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		panic(err)
	}
	CleanUpFunc := func() {
		connDb.Close()
	}
	return connDb, nil, CleanUpFunc
}
