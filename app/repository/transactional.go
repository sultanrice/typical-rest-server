package repository

import (
	"context"
	"database/sql"

	"github.com/typical-go/typical-rest-server/pkg/dbkit"
	"go.uber.org/dig"
)

// Transactional database
type Transactional struct {
	dig.In
	*sql.DB
}

// CommitMe to create begin transaction and return commit function to be deffered
func (t *Transactional) CommitMe(ctx *context.Context) func() {
	var (
		tx  *sql.Tx
		err error
	)
	if tx, err = t.DB.BeginTx(*ctx, nil); err != nil {
		*ctx = dbkit.SetErrCtx(*ctx, err)
		return func() {}
	}
	*ctx = dbkit.SetTxCtx(*ctx, tx)
	return func() {
		if err = tx.Commit(); err != nil {
			*ctx = dbkit.SetErrCtx(*ctx, err)
		}
	}
}