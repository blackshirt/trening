package organisasi

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type OrgRepository interface {
	GetByID(ctx context.Context, id int) (models.Organisasi, error)
}

type mysqlOrgRepo struct {
	db *sql.DB
}

func NewMysqlOrgRepo(db *sql.DB) OrgRepository {
	return &mysqlOrgRepo{db: db}
}

func (m mysqlOrgRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.Organisasi, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	o := &models.Organisasi{}
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

func (m mysqlOrgRepo) GetByID(ctx context.Context, id int) (models.Organisasi, error) {
	query := `SELECT * FROM organisasi WHERE id=?`
	return m.getOne(ctx, query, id)
}
