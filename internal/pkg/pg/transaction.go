package pg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)


type PgTransactionFn func(pgTX pgx.Tx) error 

func HandlePgTransaction(ctx context.Context, pgTxFn PgTransactionFn) error {
	var pgTx pgx.Tx

	_, err := pgTx.Begin(ctx)
	if err != nil {
		return fmt.Errorf("Failed begin transaction: %w", err)
	}

	errFn := pgTxFn(pgTx)

	if errFn != nil {
		if err := pgTx.Rollback(ctx); err != nil {
			return fmt.Errorf("failed rollback transaction: %w", err)
		}

		return fmt.Errorf("failed executing transaction: %w", errFn)
	}

	if err := pgTx.Commit(ctx); err != nil {
		return fmt.Errorf("failed commit transaction: %w", err)
	}

	return nil

}