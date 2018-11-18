package trx

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type TrxHistoryRepo interface {
}

type trxHistory struct {
	db *sql.DB
}

func (th *trxHistory) HistoryById(ctx context.Context, id int) (*models.TrxHistory, error) {
	query := `SELECT * FROM trainix_history WHERE id=?`
	row := th.db.QueryRowContext(ctx, query, id)

	trx := new(models.TrxHistory)
	if err := row.Scan(
		&trx.ID,
		&trx.Trx,
		&trx.Start,
		&trx.Finish,
	); err != nil {
		return nil, err
	}
	return trx, nil
}
