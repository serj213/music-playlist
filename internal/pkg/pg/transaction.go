package pg

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
)


type BunTransactionFn func(tx bun.Tx) error



func HandleBunTransaction(ctx context.Context, bunTx BunTransactionFn, db *DB) (err error) {


	tx, err := db.BeginTx(ctx,nil)
	if err != nil {
		return fmt.Errorf("failed start transaction: %w", err)
	}

	err = bunTx(tx)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return fmt.Errorf("failed executing transaction: %w: failed rollbalck transaction: %w", err, errRollback)
		}

		return fmt.Errorf("failed executing transaction: %w", err)
	}
	
	if err = tx.Commit();  err != nil {
		return fmt.Errorf("failed commit transaction: %w", err)
	}

	return nil
}