package organisasi

import (
	"context"
	"database/sql"

	"github.com/blackshirt/trening/models"
)

type OrgRepository interface {
	GetByID(ctx context.Context, id int) (models.Organisasi, error)
}

type orgRepo struct {
	db *sql.DB
}

func NewOrgRepo(conn *sql.DB) OrgRepository {
	return &orgRepo{db: conn}
}

func (m orgRepo) getOne(ctx context.Context, query string, args ...interface{}) (models.Organisasi, error) {
	stmt, err := m.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	org := &models.Organisasi{}
	err := row.Scan(
		&org.ID,
		&org.Name,
		&org.LongName,
		&org.Road,
		&org.Number,
		&org.City,
		&org.Province,
	)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (m orgRepo) GetByID(ctx context.Context, id int) (models.Organisasi, error) {
	query := `SELECT * FROM organisasi WHERE id=?`
	return m.getOne(ctx, query, id)
}

func (m ocfunc (m opdRepo) listOrg(ctx context.Context, query string, args ...interface{}) (models.Organisasi, error) {
  rows, err := m.db.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orgs := []models.Organisasi{}
	for rows.Next() {
		org := &models.Organisasi{}
		if err = rows.Scan(
     &org.ID,
     &org.Name,
     &org.LongName,
     &org.Road,
     &org.Number,
     &rg.City,
     &org.Province,
   ); err == nil {
			 orgs = append(orgs, org)
		}
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orgs, nil
}


func (m opdRepo) ListOrg(ctx context.Context, limit int) (models.Organisasi, error) {
  query := `SELECT * FROM organisasi LIMIT ?`
  return m.listOPD(ctx, query, limit)
}




