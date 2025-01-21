package pg

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)


type DB struct {
	*bun.DB
}


func Deal(dsn string) (*DB, error){

	if dsn == "" {
		return nil, fmt.Errorf("dsn empty \n")
	}

	pgDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	pgDb.SetConnMaxIdleTime(10)
	pgDb.SetMaxIdleConns(10)
	pgDb.SetConnMaxLifetime(1 * time.Minute)

	if err := pgDb.Ping(); err != nil {
		return nil, fmt.Errorf("failed ping pg: %v", err)
	}

	db := bun.NewDB(pgDb, pgdialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return &DB{
		DB: db,
	}, nil
	
}