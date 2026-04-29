package main

import (
	"context"
	"database/sql"
)

func InsertCotacao(ctx context.Context, db *sql.DB, c *Cotacao) error {
	query := `
		INSERT INTO cotacoes (
			code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, err := db.ExecContext(ctx, query,
		c.Code,
		c.Codein,
		c.Name,
		c.High,
		c.Low,
		c.VarBid,
		c.PctChange,
		c.Bid,
		c.Ask,
		c.Timestamp,
		c.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err == nil {
		c.ID = int(id)
	}

	return nil
}
