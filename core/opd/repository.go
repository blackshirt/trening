package opd

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type OPDRepository interface {
	GetByID(ctx context.Context, id int) (models.OPD, error)
}

type mysqlOPDRepo struct {
	db *sql.DB
}

func NewMysqlOPDRepo(db *sql.DB) OPDRepository {
	return &mysqlOPDRepo{db: db}
}

func (m mysqlOPDRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.OPD, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	o := &models.OPD{}
	err := row.Scan(
		&o.ID,
		&o.Name,
		&o.LongName,
		&o.Road,
		&o.Number,
		&o.City,
		&o.Province,
	)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (m mysqlOPDRepo) GetByID(ctx context.Context, id int) (models.OPD, error) {
	query := `SELECT * FROM opd WHERE id=?`
	return m.getOne(ctx, query, id)
}
