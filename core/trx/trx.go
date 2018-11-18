package trx

import (
	"context"
	"database/sql"
	"log"

	"github.com/blackshirt/trening/models"
)

type TrxRepo interface {
	TrxById(ctx context.Context, id int) (*models.Trx, error)
	TrxByName(ctx context.Context, name string) (*models.Trx, error)
	TrxList(ctx context.Context, limit, offset int) ([]*models.Trx, error)
	TrxCreate(ctx context.Context, input models.TrxInput) (*models.Trx, error)
}

type trxRepo struct {
	db *sql.DB
}

func NewTrxRepo(conn *sql.DB) TrxRepo {
	return &trxRepo{db: conn}
}

func (t *trxRepo) getOne(ctx context.Context, query string, args ...interface{}) (*models.Trx, error) {
	stmt, err := t.db.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	defer stmt.Close()

	trx := new(models.Trx)
	if err := row.Scan(
		&trx.ID,
		&trx.Name,
		&trx.Description,
		&trx.Category,
		&trx.Type,
	); err != nil {
		log.Fatal(err)
	}

	return trx, nil
}

func (t *trxRepo) TrxById(ctx context.Context, id int) (*models.Trx, error) {
	query := `SELECT * FROM trx_master WHERE id=?`
	return t.getOne(ctx, query, id)
}

func (t *trxRepo) TrxByName(ctx context.Context, name string) (*models.Trx, error) {
	query := `SELECT * FROM trx_master WHERE name=?`
	return t.getOne(ctx, query, name)
}

func (t *trxRepo) exists(ctx context.Context, name string) bool {
	query := `SELECT name FROM trx_master WHERE name=?`
	var trxname string
	err := t.db.QueryRowContext(ctx, query, name).Scan(&trxname)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No trx with name: %s", name)
		return false
	case err != nil:
		return false
	default:
		log.Printf("There is trx name in db. that is = %s", trxname)
		return true
	}
}

func (t *trxRepo) listTrx(ctx context.Context, query string, args ...interface{}) ([]*models.Trx, error) {
	rows, err := t.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	trxs := make([]*models.Trx, 0)
	for rows.Next() {
		trx := new(models.Trx)
		if err = rows.Scan(
			&trx.ID,
			&trx.Name,
			&trx.Description,
			&trx.Category,
			&trx.Type,
		); err == nil {
			trxs = append(trxs, trx)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return trxs, nil
}

func (t *trxRepo) TrxList(ctx context.Context, limit, offset int) ([]*models.Trx, error) {
	query := `SELECT * FROM trx_master LIMIT ? OFFSET ?`
	return t.listTrx(ctx, query, limit, offset)
}

func (t *trxRepo) TrxCreate(ctx context.Context, input models.TrxInput) (*models.Trx, error) {
	exist := t.exists(ctx, input.Name)
	if !exist {
		query := `INSERT INTO trx_master(name, description) VALUES(?,?)`
		_, err := t.db.ExecContext(ctx, query, input.Name, input.Description)
		if err != nil {
			return nil, err
		}
	}

	row, err := t.TrxByName(ctx, input.Name)
	if err != nil {
		return nil, err
	}
	return row, nil
}
