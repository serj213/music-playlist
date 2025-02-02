package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


type PgTransactionFn func(pgTX pgx.Tx) error 

func HandlePgTransaction(ctx context.Context, pgTxFn PgTransactionFn, db *DB) error {
	var pgTx pgx.Tx

	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("Failed begin transaction: %w", err)
	}

	errFn := pgTxFn(pgTx)

	if errFn != nil {
		if err := tx.Rollback(ctx); err != nil {
			return fmt.Errorf("failed rollback transaction: %w", err)
		}

		return fmt.Errorf("failed executing transaction: %w", errFn)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed commit transaction: %w", err)
	}

	return nil

}