package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


type DB struct {
	*pgx.Conn
}


func Deal(dsn string) (*DB, error){

	if dsn == "" {
		return nil, fmt.Errorf("dsn empty \n")
	}

	conn, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, fmt.Errorf("failed connect database postgres: %w", err)
	}

	return &DB{Conn: conn}, nil
}