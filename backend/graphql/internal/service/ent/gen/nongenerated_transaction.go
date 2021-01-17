package gen

import (
	"context"
	"database/sql/driver"
	"errors"
)

// OpenTx opens a transaction and returns a transactional
// context along with the created transaction.
func (c *Client) OpenTx(ctx context.Context) (context.Context, driver.Tx, error) {
	tx, err := c.Tx(ctx)
	if err != nil {
		return nil, nil, err
	}
	ctx = NewTxContext(ctx, tx)
	ctx = NewContext(ctx, tx.Client())
	return ctx, tx, nil
}

// OpenTxFromContext open transactions from client stored in context.
func OpenTxFromContext(ctx context.Context) (context.Context, driver.Tx, error) {
	client := FromContext(ctx)
	if client == nil {
		return nil, nil, errors.New("no client attached to context")
	}
	return client.OpenTx(ctx)
}
